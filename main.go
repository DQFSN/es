package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mq/es"
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

			params := make(map[string]interface{})
			//对应需修改的ES中文档的ID
			params["docID"] = ctx.Query("docID")

			params["name"] = ctx.Query("name")
			params["deleted"] = ctx.Query("deleted")
			params["weight"] = ctx.Query("weight")
			params["wiki"] = ctx.Query("wiki")
			params["assess_dirs"] = ctx.Query("assess_dirs")
			params["teaching_objective"] = ctx.Query("teaching_objective")
			params["desc"] = ctx.Query("desc")
			params["keywords"] = ctx.Query("keywords")
			params["desc_tex"] = ctx.Query("desc_tex")
			params["extra_desc_tex"] = ctx.Query("extra_desc_tex")
			idMap := make(map[string]interface{})
			idMap["$oid"] = ctx.Param("id")
			params["id"] = idMap

			var es es.ES
			_ = es.NewClient()

			result := es.UpdateKtag(params)
			ctx.String(200,"%s",result)
		})

		v1.PUT("/api/ktag", func(ctx *gin.Context) {

			/*
			"_id": null,
				"name": "测试知识点的子节点",
				"type": 1102,
				"weight": 0,
				"wiki": "单身的",
				"desc": "",
				"desc_tex": "",
				"extra_desc": "",
				"extra_desc_tex": "",
				"deleted": false,
				"assess_dirs": [1, 2],
				"path": "知识点 > 知识点的子节点",
				"parent_id": "50e227000000000000001102",
				"teaching_objective": {
					"desc": "单身的",
					"level": "A",
					"lesson_count": 10
				},
				"stats": {
					"freq": 0,
					"avg_diff": 0
				},
				"keywords": ["测试"]
			 */

			params := make(map[string]interface{})


			params["name"] = ctx.Query("name")
			params["type"] = ctx.Query("type")
			params["weight"] = ctx.Query("weight")
			params["wiki"] = ctx.Query("wiki")
			params["desc"] = ctx.Query("desc")
			params["desc_tex"] = ctx.Query("desc_tex")
			params["extra_desc"] = ctx.Query("extra_desc")
			params["extra_desc_tex"] = ctx.Query("extra_desc_tex")
			params["deleted"] = ctx.Query("deleted")
			params["assess_dirs"] = ctx.Query("assess_dirs")
			params["path"] = ctx.Query("path")
			params["parent_id"] = ctx.Query("parent_id")
			params["keywords"] = ctx.Query("keywords")
			object := make(map[string]interface{})
			object["freq"] = ctx.Query("freq")
			object["avg_diff"] = ctx.Query("avg_diff")
			params["stats"] = object
			object = make(map[string]interface{})
			object["desc"] = ctx.Query("extra_desc_tex")
			object["level"] = ctx.Query("extra_desc_tex")
			object["lesson_count"] = ctx.Query("extra_desc_tex")
			params["teaching_objective"] = object



			var es es.ES
			_ = es.NewClient()

			result := es.UpdateKtag(params)
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
