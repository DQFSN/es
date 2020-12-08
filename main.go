package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"mq/es"
	"mq/model"
)

func main() {

	router := gin.Default()

	v1 := router.Group("/v1")
	{

		v1.GET("/api/ktag", func(ctx *gin.Context) {
			var es es.ES
			es.NewClient()
			result := es.QueryKtag()
			ctx.String(200,"%s",result)
		})

		v1.GET("/api/ktag/:id", func(ctx *gin.Context) {
			var es es.ES
			es.NewClient()
			result := es.QueryKtagByID(ctx.Param("id"))
			ctx.String(200,"%s",result)
		})

		v1.PATCH("/api/ktag/:id", func(ctx *gin.Context) {

			var ktag model.Ktag

			if err := ctx.BindJSON(&ktag) ; err ==  nil {
				ktag.Id = ctx.Param("id")
			}else {
				log.Println(err)
			}

			var es es.ES
			_ = es.NewClient()

			result := es.UpdateKtag(ktag)
			ctx.String(200,"%s",result)
		})

		v1.PUT("/api/ktags", func(ctx *gin.Context) {

			var ktag model.Ktag

			if err := ctx.BindJSON(&ktag);err != nil {
				log.Printf("create ktag : ",err)
			}

			var es es.ES
			_ = es.NewClient()

			result := es.CreateKtag(ktag)
			ctx.String(200,"%s",result)
		})

		v1.GET("/api/item_search", func(ctx *gin.Context) {
			var es es.ES
			es.NewClient()

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
			result := es.QueryItem(params)
			fmt.Println(ctx.Request.URL.RawQuery)
			ctx.String(200,"%s",result)
		})

		v1.GET("/api/item/:id", func(ctx *gin.Context) {
			var es es.ES
			es.NewClient()

			itemID := ctx.Param("id")
			fmt.Println(itemID)
			result := es.QueryItemByID(itemID)
			ctx.String(200,"%s",result)
		})

		v1.GET("/api/omega_papers", func(ctx *gin.Context) {
			var es es.ES
			es.NewClient()
			//name=&year=0&grade=0&cat=&uid=5fb33927210b2863adb2734d
			params := make(map[string]interface{})
			params["name"] = ctx.Query("name")
			params["year"] = ctx.Query("year")
			params["grade"] = ctx.Query("grade")
			params["cat"] = ctx.Query("cat")
			params["uid"] = ctx.Query("uid")

			result := es.QueryPapers(params)
			ctx.String(200,"%s",result)
		})

		/*
			走QueryPapers能返回，但是数据字段不对，还没弄清楚plans到底是怎么查出的，但是
			创建教案和试卷的过程是相似的，只是教案将题目分为了三个部分（例题、习题、随堂测试）
		*/
		v1.GET("/api/plans", func(ctx *gin.Context) {
			var es es.ES
			es.NewClient()
			//owner_id=用户ID&grade=0&keywords=&remark=
			params := make(map[string]interface{})
			params["owner_id"] = ctx.Query("name")
			params["grade"] = ctx.Query("grade")
			params["uid"] = ctx.Query("uid")
			params["remark"] = ctx.Query("remark")

			result := es.QueryPapers(params)
			ctx.String(200,"%s",result)
		})



	}

	router.Run()


}
