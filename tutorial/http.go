package tutorial

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func Http() {
	code := 200
	fmt.Println(http.StatusText(code))

	url := "https://lok.me"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Printf("parsing %s as HTML: %v", url, err)
		os.Exit(1)
	}

	fmt.Println(doc)

	fmt.Println("done")

}
