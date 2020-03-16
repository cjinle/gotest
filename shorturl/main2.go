package main

import (
	"io"
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "index page")
}

func main() {
	http.HandleFunc("/", IndexHandler)

	log.Fatal(http.ListenAndServe(":9999", nil))
}
