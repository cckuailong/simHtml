package example

import (
	"fmt"
	"simHtml/simHtml"
)

func test(){
	f1 := "example/data/exp1.html"
	f2 := "example/data/exp2.html"
	fmt.Println(simHtml.GetSimRate(f1, f2))
}