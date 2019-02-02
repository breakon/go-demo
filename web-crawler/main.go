package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	ExampleScrape()
}
func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get("https://news.ycombinator.com/") //原生
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status) //停止并且退出程序
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	// bb := doc.Find(".athing .title a").Eq(0) 获取第一个
	// var slc1 []int = arr[:]
	doc.Find(".itemlist .athing").Find(".votelinks").Each(func(i int, selection *goquery.Selection) {
		val := selection.Next().Find("a")
		var text = val.Text()
		// var href = val.Attr("href")
		fmt.Println(i+1, text)
		// fmt.Println(href)

	})
}
