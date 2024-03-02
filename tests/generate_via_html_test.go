package tests

import (
	"testing"
	"time"

	"github.com/CirkaN/go-pdf/internal"
)

func TestGenerateFromUrl(t *testing.T) {
	saveToPath := "../storage/"
	fileName := time.Now().Format("01-02 15:04:05") + ".pdf"
	waitTime := 1 * time.Second
	url := "https://google.com"

	err := internal.GenerateFromUrl(saveToPath, fileName, url, waitTime)
	if err != nil {
		t.Errorf("Error generating PDF from URL: %v", err)
	}
}
func TestGenerateFromPdf(t *testing.T) {
	html := "../html_templates/test1.html"
	saveToPath := "../storage/"
	fileName := time.Now().Format("01-02 15:04:05") + ".pdf"
	waitTime := 1 * time.Second
	err:=internal.GenerateFromHtml(saveToPath, fileName, html, waitTime)

	if err != nil {
		t.Errorf("Error generating PDF from HTML: %v", err)
	}
}