package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func main() {

	doc, err := goquery.NewDocument("https://event.cloudnativedays.jp/cndt2022/talks?utm_source=twregular2210&utm_medium=twitter&utm_campaign=CNDT2022CFP")
	if err != nil {
		panic("htmlの取得に失敗しました")
	}

	var cnt int
	title := doc.Find("body#page-top > div#wrapper > div.container > div.row > table.table > tbody > tr")
	title.Each(func(i int, s *goquery.Selection) {
		author := s.Find(fmt.Sprintf("td.td-class-%d > ul", i+1))
		author.Each(func(i int, s *goquery.Selection) {
			cnt++
			fmt.Println(standardizeSpaces(s.Text()))
		})
	})
	fmt.Println("合計:", cnt, " 件")
}
