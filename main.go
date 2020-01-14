package main

import (
	"fmt"
	"simHtml/simHtml"
)

func main(){
	url1 := "http://lovebear.top/2020/01/08/PyBgpStream_Install/"
	url2 := "http://lovebear.top/2020/01/06/Grafana_Install_And_Config/"
	fmt.Println(simHtml.GetSimFromUrl(url1, url2))
}
