package main

import (
	"time"

	"github.com/CirkaN/go-pdf/internal"
)

func testReportFromHtml(){
	html := "../../html_templates/test1.html"
	saveToPath := "../../storage/"
	fileName := time.Now().Format("01-02 15:04:05") + ".pdf"
	//raise for chart loading etc..
	waitTime := 1 * time.Second
	internal.GenerateFromHtml(saveToPath, fileName, html, waitTime)
}

func testReportFromUrl(){
	saveToPath := "../../storage/"
	fileName := time.Now().Format("01-02 15:04:05") + ".pdf"
	waitTime := 15 * time.Second
	url := "https://moj-biznis.rs"
	internal.GenerateFromUrl(saveToPath, fileName, url, waitTime)
}
func main() {
	testReportFromHtml()
}
