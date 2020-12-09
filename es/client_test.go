package es

import (
	"fmt"
	"testing"
)

func Test_Client(t *testing.T) {
	var es ES
	_ = es.NewClient()
	params := make(map[string]interface{})
	fmt.Println(es.QueryKtag(params))

	params["difficulty"] = 1
	params["keywords"] = ""
	params["item_type"] = 1000
	params["edu"] = 2
	params["classes"] = ""
	params["tag_ids"] = ""
	fmt.Println(es.QueryItem(params))

	fmt.Println(es.QueryItemByID(""))
	fmt.Println(es.QueryItemByID("dadsadas"))
	fmt.Println(es.QueryItemByID("50e227000000000000001103"))
}
