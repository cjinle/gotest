package main

import (
	"github.com/gocolly/colly"
	"fmt"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("news.163.com"),
	)
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnRequest(func(r *colly.Request){
		fmt.Println("visiting", r.URL.String())
	})
	c.Visit("https://news.163.com/")
	fmt.Println(c)
}