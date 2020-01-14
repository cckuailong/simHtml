package example

import (
	"fmt"
	"github.com/cckuailong/simHtml/simHtml"
)

func testFileInterface(){
	f1 := "example/data/exp1.html"
	f2 := "example/data/exp2.html"
	fmt.Println(simHtml.GetSimFromFile(f1, f2))
}

func testStrInterface(){
	html1 := "<html><head></head></html>"
	html2 := "<html><body></body></html>"
	fmt.Println(simHtml.GetSimFromStr(html1, html2))
}

func testUrlInterface(){
	url1 := "http://lovebear.top/2020/01/08/PyBgpStream_Install/"
	url2 := "http://lovebear.top/2020/01/06/Grafana_Install_And_Config/"
	fmt.Println(simHtml.GetSimFromUrl(url1, url2))
}
