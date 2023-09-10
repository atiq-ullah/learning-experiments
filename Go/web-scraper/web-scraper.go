package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func main() {
	// Create a new collector
	c := colly.NewCollector()

	// On every <a> element with href attribute call this callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// Visit the link found on the page
		// e.Request.Visit(e.Attr("href"))  // Uncomment this if you want to follow links

		// Print the link
		fmt.Println(e.Attr("href"))
	})

	// On every HTML element which has <title> tag call this
	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println("Page title: ", e.Text)
	})

	// Log errors
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	// Start the web scraping by visiting a website
	err := c.Visit("http://google.com")
	if err != nil {
		log.Fatal(err)
	}
}
