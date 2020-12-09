package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var defaultURI = "mongodb://localhost:27017"
type MyClient struct {
	URI string
	Client *mongo.Client
}


func (c *MyClient) Connect()  {

	if c.Client == nil  {
		if len(c.URI) == 0 {
			c.URI = defaultURI
		}
		client,err := mongo.NewClient(options.Client().ApplyURI(c.URI))
		checkErr(err)
		c.Client = client
	}

	ctx, cancel := context.WithTimeout(context.Background(),time.Second*10)
	defer cancel()

	err := c.Client.Connect(ctx)
	checkErr(err)

	err = c.Client.Ping(ctx,readpref.Primary())
	if err == nil {
		fmt.Println("连接数据库成功～")
	}else {
		fmt.Println("连接数据库失败～")
	}

}

func (c *MyClient) Disconnect() {
	err := c.Client.Disconnect(context.Background())
	if err != nil {
		fmt.Println("数据库关闭失败～")
	}else {
		fmt.Println("数据库关闭成功～")
	}
}

func (c *MyClient) Query(coll string) (result []*map[string]interface{}) {
	//collection := c.Client.Database("bison").Collection(coll)
	collection := c.Client.Database("ir").Collection(coll)
	ctx, cancel := context.WithTimeout(context.Background(),time.Second*10)
	defer cancel()

	limit := int64(0)
	cur,err := collection.Find(ctx,bson.D{},&options.FindOptions{Limit: &limit})
	checkErr(err)

	if err :=cur.All(context.Background(),&result);err != nil {
		fmt.Println(err)
	}
	//for cur.Next(context.Background()) {
	//	record := fmt.Sprintf("%s",cur.Current)
	//	result = append(result,strings.Replace(record,"_id","id",1))
	//}

	return
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}