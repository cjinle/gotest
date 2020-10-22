package main

import (
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	// outputDir := "data/"
	c := colly.NewCollector()
	c.OnHTML("html", func(e *colly.HTMLElement) {
		log.Println("OnHTML: ", e.Request.URL)
		// log.Println("index: ", strings.Index(e.Request.URL.String(), "/VOA_Standard_English/"))
		if strings.Index(e.Request.URL.String(), "/VOA_Standard_English/") > -1 {
			// log.Println("title: ", e.DOM.Find("h1").Text())
			log.Println("title: ", e.ChildText("h1"))
			log.Println("content: ", e.ChildText("div.Content"))
		}
		e.ForEach("div.List li a", func(_ int, el *colly.HTMLElement) {
			log.Println("page href: ", el.Attr("href"))
			c.Visit(el.Request.AbsoluteURL(el.Attr("href")))
		})
		e.ForEach("div.pagelist a", func(_ int, el *colly.HTMLElement) {
			// log.Println("page href: ", el.Attr("href"))
			c.Visit(el.Request.AbsoluteURL(el.Attr("href")))
		})

	})


	c.OnError(func(r *colly.Response, e error) {
		log.Println("error:", e, r.Request.URL, string(r.Body))
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println("OnResponse: ", r.Request.URL)

		// if strings.Index(r.Headers.Get("Content-Type"), "image") > -1 {
		// 	r.Save(outputDir + r.FileName())
		// 	return
		// }
		return

	})

	c.Visit("https://www.51voa.com/VOA_Standard_1.html")
}
