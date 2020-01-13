package simHtml

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"strings"
)

func GetDomCssList(filePath string) ([]string, []string){
	queue := []*goquery.Selection{}
	dom_res, css_res := []string{}, []string{}
	exp1,_ := ioutil.ReadFile(filePath)
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(exp1))
	queue = append(queue, doc.Find("html"))
	for len(queue) > 0{
		cur_sel := queue[0]
		queue = queue[1:]
		if len(cur_sel.Nodes)==0{
			continue
		}
		for _, c := range(cur_sel.Nodes){
			dom_res = append(dom_res, c.Data)
			for _,item := range(c.Attr){
				key := strings.ToLower(item.Key)
				if key == "class" || key == "style"{
					css_res = append(css_res, item.Val)
				}
			}
		}
		queue = append(queue, cur_sel.Children())
	}
	return dom_res, css_res
}

func GetSimRate(file1,file2 string) float64{
	domList1, cssList1 := GetDomCssList(file1)
	domList2, cssList2 := GetDomCssList(file2)
	domSimNum := LongestCommonSubsequence(domList1, domList2)
	cssSimNum := LongestCommonSubsequence(cssList1, cssList2)
	domRate := float64(2*domSimNum)/float64(len(domList1)+len(domList2))
	cssRate := float64(2*cssSimNum)/float64(len(cssList1)+len(cssList2))
	return 0.3*domRate + 0.7*cssRate
}
