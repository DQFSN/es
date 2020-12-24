package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"mq/es"
	"mq/model"
)

func	getKtags (ctx *gin.Context) {
	params := make(map[string]interface{})
	params["type"] = ctx.Query("type")
	var client es.ES
	_ = client.NewClient()
	result := client.QueryKtag(params)
	ctx.String(200,"%s",result)
}

func	getKtagByID (ctx *gin.Context) {
	var client es.ES
	_ = client.NewClient()
	result := client.QueryKtagByID(ctx.Param("id"))
	ctx.String(200,"%s",result)
}

func	updateKtagByID(ctx *gin.Context) {

	var ktag model.Ktag

	if err := ctx.BindJSON(&ktag) ; err ==  nil {
		ktag.Id = ctx.Param("id")
	}else {
		log.Println(err)
	}

	var client es.ES
	_ = client.NewClient()

	result := client.UpdateKtag(ktag)
	ctx.String(200,"%s",result)
}

func	createKtag (ctx *gin.Context) {
	var ktag model.Ktag

	if err := ctx.BindJSON(&ktag);err != nil {
		log.Printf("create ktag : %v",err)
	}

	var client es.ES
	_ = client.NewClient()

	result := client.CreateKtag(ktag)
	ctx.String(200,"%s",result)
}

func 	getItemIDBySearch (ctx *gin.Context) {
	var client es.ES
	_ = client.NewClient()

	params := make(map[string]interface{})
	params["difficulty"] = ctx.Query("difficulty")
	params["keywords"] = ctx.Query("keywords")
	params["edu"] = ctx.Query("edu")
	params["item_type"] = ctx.Query("item_type")
	params["classes"] = ctx.Query("classes")
	params["tag_ids"] = ctx.Query("tag_ids")
	params["keywords"] = ctx.Query("keywords")
	params["page_size"] = ctx.Query("page_size")
	params["page_num"] = ctx.Query("page_num")
	result := client.QueryItem(params)
	ctx.String(200,"%s",result)
}

func	getItemByID (ctx *gin.Context) {
	var client es.ES
	_ = client.NewClient()

	itemID := ctx.Param("id")
	result := client.QueryItemByID(itemID)
	ctx.String(200,"%s",result)
}

func getOmegaPapers (ctx *gin.Context) {
	var client es.ES
	_ = client.NewClient()
	//name=&year=0&grade=0&cat=&uid=5fb33927210b2863adb2734d
	params := make(map[string]interface{})
	params["name"] = ctx.Query("name")
	params["year"] = ctx.Query("year")
	params["grade"] = ctx.Query("grade")
	params["cat"] = ctx.Query("cat")
	params["uid"] = ctx.Query("uid")

	result := client.QueryOmegaPapers(params)
	ctx.String(200,"%s",result)
}

func	getPlans (ctx *gin.Context) {
	var client es.ES
	_ = client.NewClient()
	//owner_id=用户ID&grade=0&keywords=&remark=
	params := make(map[string]interface{})
	params["owner_id"] = ctx.Query("name")
	params["grade"] = ctx.Query("grade")
	params["uid"] = ctx.Query("uid")
	params["remark"] = ctx.Query("remark")

	result := client.QueryOmegaPapers(params)
	ctx.String(200,"%s",result)
}

func getPapers(ctx *gin.Context) {
	//owner_id=5fb33927210b2863adb2734d&grade=0&keywords=&remark=
	params := make(map[string]interface{})
	params["owner_id"] = ctx.Query("owner_id")
	params["grade"] = ctx.Query("grade")
	params["keywords"] = ctx.Query("keywords")
	params["remark"] = ctx.Query("remark")

	var client es.ES
	_ = client.NewClient()

	result := client.GetPapers(params)
	ctx.String(200,result)
}

func main() {

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		//题目
		v1.GET("/api/item_search", getItemIDBySearch)

		v1.GET("/api/item/:id", getItemByID)

		//知识点
		v1.PUT("/api/ktags", createKtag)

		v1.GET("/api/ktags", getKtags)

		v1.GET("/api/ktag/:id", getKtagByID)

		v1.PATCH("/api/ktag/:id", updateKtagByID)

		//教案
		/*
			走QueryPapers能返回，但是数据字段不对，还没弄清楚plans到底是怎么查出的，但是
			创建教案和试卷的过程是相似的，只是教案将题目分为了三个部分（例题、习题、随堂测试）
		*/
		v1.GET("/api/plans", getPlans)

		//试卷
		v1.GET("/api/papers", getPapers)

		v1.GET("/api/omega_papers", getOmegaPapers)


	}

	router.Run()


}
