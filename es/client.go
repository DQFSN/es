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

func (es *ES) QueryKtag(params map[string]interface{}) string {

	queryStr := ""
	if params["type"] != "" {
		bodyBytes,_ := json.Marshal(params)
		queryStr = `_source:"`+string(bodyBytes)+`"`
	}

	log.Printf("--------------------QueryKtag-------------------\n %v\n%v\n",params,queryStr)
	//es.Client.m
	rsp, err := es.Client.Search(func(request *esapi.SearchRequest) {
		request.Index = []string{"ktag"}
		request.Pretty = true
		request.Query = queryStr
		request.Source = []string{"_id","weight","deleted","keywords","name","parent_id","stats","type"}
	})
	if err != nil {
		log.Println(err)
	}
	defer rsp.Body.Close()
	buffer,_ := ioutil.ReadAll(rsp.Body)
	return string(buffer)
}

func (es *ES) QueryKtagByID(id string) string {

	log.Printf("--------------------QueryKtagByID-------------------\n%v\n",id)

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

	log.Printf("--------------------UpdateKtag-------------------\n%v\n%v\n",ktag,string(updateBodyByte))

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

	log.Printf("--------------------CreateKtag-------------------\n%v\n%v\n",ktag,string(bodyBytes))


	ctx, cancel := context.WithTimeout(context.Background(),time.Second*10)
	defer cancel()
	rsp, err := req.Do(ctx,es.Client)
	checkErr(err)
	defer rsp.Body.Close()

	buffer,_ := ioutil.ReadAll(rsp.Body)

	return string(buffer)

}

func (es *ES) QueryItem(params map[string]interface{}) string {
	queryStr := ""

	for k,v := range params {
		if v == "" || v == 0 {
			delete(params,k)
		}
	}

	if len(params) > 0 {
		bodyBytes, _ := json.Marshal(params)
		queryStr = `_source:"`+string(bodyBytes)+`"`
	}

	log.Printf("--------------------QueryItem-------------------\n%v\n%v\n",params,queryStr)

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

	queryStr := "_source:\"{id:\""+id+"\"}\""

	log.Printf("--------------------QueryItemByID-------------------\n%v\n",queryStr)

	rsp, err := es.Client.Search(func(request *esapi.SearchRequest) {
		request.Index = []string{"item"}
		request.Query = queryStr
		request.Pretty = true
	})
	if err != nil {
		log.Println(err)
	}
	buffer, err := ioutil.ReadAll(rsp.Body)
	return string(buffer)
}

func (es *ES) QueryOmegaPapers(params map[string]interface{}) string {

	queryStr := ""

	for k,v := range params {
		if v == "" || v == 0 {
			delete(params,k)
		}
	}

	if len(params) > 0 {
		bodyBytes, _ := json.Marshal(params)
		queryStr = `_source:"`+string(bodyBytes)+`"`
	}

	log.Printf("--------------------QueryOmegaPapers-------------------\n%v\n%v\n",params,queryStr)

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

func (es *ES) GetPapers(params map[string]interface{}) string {
	queryStr := ""

	for k, v := range params {
		if v == "" {
			delete(params,k)
		}
	}

	if len(params) > 0 {
		bodyBytes,_ := json.Marshal(params)
		queryStr = `_source:"`+string(bodyBytes)+`"`
	}

	log.Printf("--------------------GetPapers-------------------\n%v\n%v\n",params,queryStr)

	rsp,_ := es.Client.Search(func(request *esapi.SearchRequest) {
		request.Index = []string{"paper"}
		request.Pretty = true
		request.Query = queryStr
	})
	defer rsp.Body.Close()
	buffer,_ := ioutil.ReadAll(rsp.Body)

	return string(buffer)

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