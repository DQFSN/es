package spider

import (
	"testing"
)

func Test_GetHTML(t *testing.T) {
		gs := GetGoods("vue")
		for i, g := range gs {
			t.Logf("%d----name:%v\n\tprice:%v\n\timgsrc:%v\n",i,g.Name,g.Price,g.ImgLink)
		}
}

