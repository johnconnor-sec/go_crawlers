package scraper

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func Scrape(doc *goquery.Document) (string, string, error) {
	title := strings.TrimSpace(doc.Find("title").Text())
	content := strings.TrimSpace(doc.Find("body").Text())
	return title, content, nil
}