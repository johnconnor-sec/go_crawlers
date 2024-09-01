package tests

import (
	"strings"
	"testing"
	"github.com/PuerkitoBio/goquery"
	"github.com/johnconnor-sec/firecrawl/internal/scraper"
)

func TestScrape(t *testing.T) {
	html := `<html><head><title>Test</title></head><body><p>Hello World</p></body></html>`
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	title, content, err := scraper.Scrape(doc)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if title != "Test" {
		t.Fatalf("Expected title 'Test', got %v", title)
	}

	if content != "Hello World" {
		t.Fatalf("Expected content 'Hello World', got %v", content)
	}
}