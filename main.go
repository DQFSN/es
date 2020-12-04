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
		v1.GET("/api/ktags", func(ctx *gin.Context) {
			var es es.ES
			es.NewClient()
			result := es.QueryKtag()
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
	}

	router.Run()


}
