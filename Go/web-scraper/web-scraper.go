package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML("a[data-track]", func(e *colly.HTMLElement) {
		e.DOM.Find("span.p--small").Each(func(index int, item *goquery.Selection) {
			fmt.Print(strings.TrimSpace(item.Text()))
			fmt.Print(" | ")
		})

		e.DOM.Find("span.smaller").Each(func(index int, item *goquery.Selection) {
			fmt.Println(strings.TrimSpace(item.Text()))
		})
	})
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	err := c.Visit("https://www.rottentomatoes.com/browse/movies_in_theaters/sort:popular")
	if err != nil {
		log.Fatal(err)
	}
}
