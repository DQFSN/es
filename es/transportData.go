package es

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	mg "mq/mongo"
	"strings"
)

func moveDataFromMGToES() {
	es, err := elasticsearch.NewDefaultClient()
	checkErr(err)

	var mgclient mg.MyClient
	mgclient.Connect()
	defer mgclient.Disconnect()

	colls := []string{"circle","cluster","hydra","item","item_tag","ktag","omega_paper","op_log","super_cluster","user"}

	for _, coll := range colls {

		queryData := mgclient.Query(coll)
		run := make(chan int, 20)
		for _, record := range queryData {
			recordByte ,_ := json.Marshal(record)
			recordStr := strings.Replace(string(recordByte),"_id","id",1)
			req := esapi.IndexRequest{
				Index: coll,
				DocumentID: fmt.Sprintf("%v",(*record)["_id"].(primitive.ObjectID).Hex()),
				Body: strings.NewReader(recordStr),
			}

			run <- 1
			go func(){
				rep, err := req.Do(context.Background(),es)
				checkErr(err)
				defer rep.Body.Close()
				fmt.Println(rep)
				<-run
			}()
		}

		fmt.Printf("%v数据量:%v",coll,len(queryData))

	}



}

func checkErr(err error) {
	if err != nil {
		log.Fatalf("ERROR: %s",err)
	}
}


//es, err := elasticsearch.NewDefaultClient()
//checkErr(err)
//log.Println(elasticsearch.Version)
//log.Println(es.Info())
//res,err := es.Index(
//	"test",
//	strings.NewReader(`{"name":"段其沣111","age":23}`),
//	es.Index.WithDocumentID("2"),
//	es.Index.WithRefresh("true"),
//	)
//checkErr(err)
//defer res.Body.Close()
//log.Println(res)

//增
//req := esapi.IndexRequest{
//	Index: 		"test",
//	Body:		strings.NewReader(`{"name":"雷新","age":25}`),
//	DocumentID: "3",
//	Refresh: 	"true",
//}

//查
//req = esapi.GetRequest{
//	Index: "test",
//	DocumentID: "1",
//}

//改
//req := esapi.UpdateRequest{
//	Index: "test",
//	DocumentID: "1",
//	Body:		strings.NewReader(`{
//			"doc":{
//				"name":"段其沣"
//			}
//			}`),
//}

//删
//req := esapi.DeleteRequest{
//Index: "test",
//DocumentID: "1",
//}
//
//
//res, err := req.Do(context.Background(), es)
