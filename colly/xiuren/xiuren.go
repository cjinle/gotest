package main

import (
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	outputDir := "data/"
	c := colly.NewCollector(
		// colly.AllowedDomains("nvshens.org"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("referer", "https://www.nvshens.org/")
	})
	c.OnHTML("html", func(e *colly.HTMLElement) {
		log.Println("OnHTML: ", e.Request.URL)
		e.ForEach("ul#hgallery img", func(_ int, el *colly.HTMLElement) {
			log.Println("img src: ", el.Attr("src"))
			c.Visit(el.Attr("src"))
		})
		e.ForEach("div#pages a", func(_ int, el *colly.HTMLElement) {
			log.Println("page href: ", el.Attr("href"))
			c.Visit(el.Request.AbsoluteURL(el.Attr("href")))
		})

	})

	c.OnError(func(r *colly.Response, e error) {
		log.Println("error:", e, r.Request.URL, string(r.Body))
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println("OnResponse: ", r.Request.URL)
		if strings.Index(r.Headers.Get("Content-Type"), "image") > -1 {
			r.Save(outputDir + r.FileName())
			return
		}
		return

	})

	c.Visit("https://www.nvshens.org/g/33128/1.html")
}
