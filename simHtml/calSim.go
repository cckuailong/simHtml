package simHtml

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func getDomCssList(doc *goquery.Document) ([]string, []string){
	queue := []*goquery.Selection{}
	dom_res, css_res := []string{}, []string{}
	queue = append(queue, doc.Selection)
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
	return dom_res[1:], css_res
}

func getSimRate(doc1, doc2 *goquery.Document) float64{
	var domRate, cssRate float64
	domList1, cssList1 := getDomCssList(doc1)
	domList2, cssList2 := getDomCssList(doc2)
	domSimNum := LongestCommonSubsequence(domList1, domList2)
	cssSimNum := LongestCommonSubsequence(cssList1, cssList2)
	domLen := len(domList1)+len(domList2)
	cssLen := len(cssList1)+len(cssList2)
	if domLen == 0{
		domRate = 0
	}else{
		domRate = float64(2*domSimNum)/float64(domLen)
	}
	if cssLen == 0{
		cssRate = 0
	}else{
		cssRate = float64(2*cssSimNum)/float64(cssLen)
	}
	return 0.3*domRate + 0.7*cssRate
}
