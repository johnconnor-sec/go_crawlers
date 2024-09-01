package tests

import (
	"testing"
	"github.com/johnconnor-sec/firecrawl/internal/crawler"
	"github.com/johnconnor-sec/firecrawl/pkg/client"
	"github.com/johnconnor-sec/firecrawl/pkg/config"
)

func TestCrawl(t *testing.T) {
	cfg := config.New(2)
	client := client.New(cfg)
	c := crawler.New(client, cfg)
	
	results, err := c.Crawl("http://example.com")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(results) == 0 {
		t.Fatalf("Expected results, got 0")
	}
}
