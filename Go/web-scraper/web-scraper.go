package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
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
		// fmt.Println(e.Attr("href"))
	})

	// On every HTML element which has <title> tag call this
	// c.OnHTML("title", func(e *colly.HTMLElement) {
	// 	fmt.Println("Page title: ", e.Text)
	// })

	c.OnHTML("a[data-track]", func(e *colly.HTMLElement) {
		// Additional search within this span
		e.DOM.Find("span.p--small").Each(func(index int, item *goquery.Selection) {
			// Do something with each "anotherSelector" within this span
			fmt.Print(strings.TrimSpace(item.Text()))
			fmt.Print(" | ")
		})

		// Additional search within this span
		e.DOM.Find("span.smaller").Each(func(index int, item *goquery.Selection) {
			// Do something with each "anotherSelector" within this span
			fmt.Println(strings.TrimSpace(item.Text()))
		})
	})

	// Log errors
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	// Start the web scraping by visiting a website
	// err := c.Visit("http://google.com")
	err := c.Visit("https://www.rottentomatoes.com/browse/movies_in_theaters/sort:popular")
	if err != nil {
		log.Fatal(err)
	}
}
