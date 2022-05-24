package main

import (
	"fmt"
	"log"
	"time"

	readability "github.com/go-shiori/go-readability"
)

var (
	urls = []string{
		// Both will be parse-able now
		"https://e360.yale.edu/digest/bugs-are-evolving-to-eat-plastic-study-finds",
	}
)

func main() {
	for _, url := range urls {
		article, err := readability.FromURL(url, 30*time.Second)

		if err != nil {
			log.Fatalf("failed to parse %s: %v\n", url, err)
		}

		fmt.Printf("URL     : %s\n", url)
		fmt.Printf("Title   : %s\n", article.Title)
		fmt.Printf("Author  : %s\n", article.Byline)
		fmt.Printf("Length  : %d\n", article.Length)
		fmt.Printf("SiteName: %s\n", article.SiteName)
		fmt.Printf("Image   : %s\n", article.Image)
		fmt.Printf("Excerpt : %s\n", article.Excerpt)
		fmt.Printf("TextContent : %s\n", article.TextContent)
		fmt.Println()
	}
}
