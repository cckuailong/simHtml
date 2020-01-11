package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
)


func main(){
	//res := simHtml.LongestCommonSubsequence([]string{"aaa","bbb","ccc","dddd"}, []string{"bbb","dddd"})
	exp1,_ := ioutil.ReadFile("example/data/exp1.html")
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(exp1))
	doc.Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.)
	})
}
