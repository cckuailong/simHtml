package simHtml

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"strings"
)

// File interface
func GetSimFromFile(file1, file2 string) float64{
	cont1, _ := ioutil.ReadFile(file1)
	doc1, _ := goquery.NewDocumentFromReader(bytes.NewReader(cont1))
	cont2, _ := ioutil.ReadFile(file2)
	doc2, _ := goquery.NewDocumentFromReader(bytes.NewReader(cont2))
	return getSimRate(doc1, doc2)
}

// String interface
func GetSimFromStr(html1, html2 string) float64{
	doc1, _ := goquery.NewDocumentFromReader(strings.NewReader(html1))
	doc2, _ := goquery.NewDocumentFromReader(strings.NewReader(html2))
	return getSimRate(doc1, doc2)
}

// Url interface
func GetSimFromUrl(url1, url2 string) float64{
	http.DefaultClient.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	resp1, err := http.Get(url1)
	if err != nil{
		fmt.Println(err)
		return 0
	}
	resp2, err := http.Get(url2)
	if err != nil{
		fmt.Println(err)
		return 0
	}
	doc1, _ := goquery.NewDocumentFromReader(resp1.Body)
	doc2, _ := goquery.NewDocumentFromReader(resp2.Body)
	return getSimRate(doc1, doc2)
}
