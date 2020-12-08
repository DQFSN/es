package es

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"log"
	"mq/model"
	"strings"
	"time"
)

type ES struct {
	Client *elasticsearch.Client
}

func (es *ES) NewClient() error {

	if es.Client == nil {
		c, err := elasticsearch.NewDefaultClient()
		if err != nil {
			return err
		}

		es.Client = c
	}

	return nil
}

func (es *ES) QueryKtag() string {

	rsp, err := es.Client.Search(func(request *esapi.SearchRequest) {
		request.Index = []string{"ktag"}
		request.Pretty = true
		size := 100
		request.Size = &size
		request.Source = []string{"_id","weight","deleted","keywords","name","parent_id","stats"}
	})
	if err != nil {
		log.Println(err)
	}
	defer rsp.Body.Close()
	buffer,_ := ioutil.ReadAll(rsp.Body)
	return string(buffer)
}

func (es *ES) QueryKtagByID(id string) string {

	rsp, err := es.Client.Search(func(request *esapi.SearchRequest) {
		request.Index = []string{"ktag"}
		request.Pretty = true
		request.Query = "_source:\"{id:\"" + id +"\"}\""
		size := 1
		request.Size = &size
	})
	if err != nil {
		log.Println(err)
	}
	defer rsp.Body.Close()
	buffer,_ := ioutil.ReadAll(rsp.Body)
	return string(buffer)
}

func (es *ES) UpdateKtag(ktag model.Ktag) string {

	ktagWithID := changeIDLabel(ktag)

	updateMap := make(map[string]interface{})
	updateMap["doc"] = ktagWithID
	updateBodyByte, _ :=json.Marshal(updateMap)

	fmt.Println(string(updateBodyByte))
	req := esapi.UpdateRequest{
		Index: "ktag",
		DocumentID: ktag.Id,
		Body: bytes.NewReader(updateBodyByte),
	}

	ctx, cancel := context.WithTimeout(context.Background(),time.Second*10)
	defer cancel()
	rsp, err := req.Do(ctx,es.Client)
	if err != nil {
		log.Println(err)
	}
	defer rsp.Body.Close()
	buffer,_ := ioutil.ReadAll(rsp.Body)
	return string(buffer)
}

func (es *ES) CreateKtag(ktag model.Ktag) string {

	id := primitive.NewObjectID()

	ktag.Id = id.Hex()
	ktagWithID := changeIDLabel(ktag)
	bodyBytes, _ :=json.Marshal(ktagWithID)

	req := esapi.CreateRequest{
		Index: "ktag",
		DocumentID: id.Hex(),
		Body: bytes.NewReader(bodyBytes),
	}

	ctx, cancel := context.WithTimeout(context.Background(),time.Second*10)
	defer cancel()
	rsp, err := req.Do(ctx,es.Client)
	checkErr(err)
	defer rsp.Body.Close()

	buffer,_ := ioutil.ReadAll(rsp.Body)

	return string(buffer)

}

func (es *ES) QueryItem(params map[string]interface{}) string {
	//开头 {
	queryStr := "_source:\"{"
	for k,v := range params {
		if k == "page_num" || k == "page_size" || v == "" {
			continue
		}
		queryStr += k+`:"`+fmt.Sprintf("%v",v)+`",`
	}
	if queryStr != "_source:\"{" {
		//去掉末尾 ','
		queryStr = queryStr[:len(queryStr)-1]
		//末尾的 }
		queryStr += "}\""
	}else {
		queryStr = ""
	}


	fmt.Println("--->",queryStr)
	rsp, err := es.Client.Search(func(request *esapi.SearchRequest) {
		request.Index = []string{"item"}
		request.Pretty = true
		request.Query = queryStr
		request.Source = []string{"_id"}
	})

	if err != nil {
		log.Println(err)
	}
	buffer ,err := ioutil.ReadAll(rsp.Body)

	result := string(buffer)
	if err != nil {
		log.Println(err)
	}

	return result
}

func (es *ES) QueryItemByID(id string) string {

	rsp, err := es.Client.Search(func(request *esapi.SearchRequest) {
		request.Index = []string{"item"}
		request.Query = "_source:\"{id:\""+id+"\"}\""
		request.Pretty = true
	})
	if err != nil {
		log.Println(err)
	}
	buffer, err := ioutil.ReadAll(rsp.Body)
	return string(buffer)
}

func (es *ES) QueryPapers(params map[string]interface{}) string {

	//开头 {
	queryStr := "_source:\"{"
	for k,v := range params {
		if v == "" {
			continue
		}
		queryStr += k+`:"`+fmt.Sprintf("%v",v)+`",`
	}
	if queryStr != "_source:\"{" {
		//去掉末尾 ','
		queryStr = queryStr[:len(queryStr)-1]
		//末尾的 }
		queryStr += "}\""
	}else {
		queryStr = ""
	}

	fmt.Println("--->",queryStr)
	rsp, err := es.Client.Search(func(request *esapi.SearchRequest) {
		request.Index = []string{"omega_paper"}
		request.Pretty = true
		request.Query = queryStr
		request.Source = []string{"_id","cat","deleted","districts","doc","edu","flow_stats","grade","item_ids","item_tpls","name","remark","uid","year"}
	})

	if err != nil {
		log.Println(err)
	}
	buffer ,err := ioutil.ReadAll(rsp.Body)

	result := string(buffer)
	if err != nil {
		log.Println(err)
	}

	return result
}

func changeIDLabel(in interface{}) map[string]interface{} {
	paramsByte,_ := json.Marshal(in)
	paramsStr := string(paramsByte)

	/*
		此处是因为，mongo原始数据中有'_id'字段，但是从mongo --> ES 时，'_id'为ES保留字段，
		顾将'_id'-->'id',但是web请求还是'_id',
		因此有了这么一步转换
	*/
	paramsStr = strings.Replace(paramsStr,"_id","id",1)
	out := make(map[string]interface{})
	_ = json.Unmarshal([]byte(paramsStr),&out)
	return out
}