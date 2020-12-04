package spider

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

func GetGoods(keyword string) (goods []Goods) {
	url := "https://search.jd.com/Search?keyword="+keyword+"&enc=utf-8"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	// 自定义Header
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_0_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.67 Safari/537.36")

	rep, err := client.Do(req)
	checkErr(err)
	defer rep.Body.Close()

	goods = make([]Goods,0)

	doc, err := goquery.NewDocumentFromReader(rep.Body)

	//content, _ := ioutil.ReadFile("html1.txt")
	//doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(content))

	doc.Find("div.gl-i-wrap").Each(func(i int, selection *goquery.Selection) {
		var name string
		var price string
		var imgSrc string
		var ok bool
		name = strings.TrimSpace(selection.Find("div.p-name").Find("a").Find("em").Text())
		price = strings.TrimSpace(selection.Find("div.p-price").Find("i").Text())
		//懒加载数据扒下来的可能性高些
		imgSrc,ok = selection.Find("div.p-img > a").Find("img").Attr("data-lazy-img")
		if !ok {
			imgSrc,ok = selection.Find(".p-img").Find("img").Attr("src")

		}

		if ok {
			var good Goods
			good.Name = name
			good.Price = price
			good.ImgLink = imgSrc
			goods = append(goods,good)
		}

	})

	return
}

func checkErr(err error) {
	if err != nil  {
		fmt.Println(err)
	}
}
