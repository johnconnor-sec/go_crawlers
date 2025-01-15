package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/johnconnor-sec/firecrawl/internal/crawler"
	"github.com/johnconnor-sec/firecrawl/internal/extractor"
	"github.com/johnconnor-sec/firecrawl/pkg/client"
	"github.com/johnconnor-sec/firecrawl/pkg/config"
)

func main() {
	url := flag.String("url", "", "URL to crawl")
	depth := flag.Int("depth", 2, "Crawl depth")
	format := flag.String("format", "markdown", "Output format: markdown or json")
	output := flag.String("output", "", "File to save the output")
	flag.Parse()

	if *url == "" {
		log.Fatal("URL is required")
	}

	cfg := config.New(*depth)
	client := client.New(cfg)
	c := crawler.New(client, cfg)
	
	results, err := c.Crawl(*url)
	if err != nil {
		log.Fatal(err)
	}

	var finalOutput string
	for _, r := range results {
		var output string
		if *format == "json" {
			jsonOutput, err := extractor.ToJSON(r.Title, r.Content, r.Paragraph, r.List, r.Header)
			if err != nil {
				log.Fatal(err)
			}
			output = string(jsonOutput)
		} else {
			output = extractor.ToMarkdown(r.Title, r.Content)
		}
		finalOutput += output + "\n"
	}

	if *output != "" {
		f, err := os.Create(*output)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		_, err = f.WriteString(finalOutput)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println(finalOutput)
	}
}
