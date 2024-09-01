package crawler

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/johnconnor-sec/firecrawl/internal/scraper"
	"github.com/johnconnor-sec/firecrawl/pkg/client"
	"github.com/johnconnor-sec/firecrawl/pkg/config"
)

type Crawler struct {
	client *client.Client
	cfg    *config.Config
}

func New(client *client.Client, cfg *config.Config) *Crawler {
	return &Crawler{
		client: client,
		cfg:    cfg,
	}
}

func (c *Crawler) Crawl(url string) ([]Result, error) {
	var results []Result
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	title, content, err := scraper.Scrape(doc)
	if err != nil {
		return nil, err
	}

	results = append(results, Result{
		URL:     url,
		Title:   title,
		Content: content,
	})

	// Additional logic for deeper crawling can be added here

	return results, nil
}

type Result struct {
	URL     string
	Title   string
	Content string
}
