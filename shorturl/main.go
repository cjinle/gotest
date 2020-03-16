package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var stat map[string]int

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v\n\n", stat)))
	w.Write([]byte(time.Now().String()))
}

func GotoHandler(w http.ResponseWriter, r *http.Request) {
	stat[r.RequestURI] += 1
	// w.Write([]byte("goto"))
	http.Redirect(w, r, "http://www.baidu.com", http.StatusSeeOther)
}

func main() {
	stat = make(map[string]int)
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/{param:[0-9a-zA-Z]+}", GotoHandler)

	log.Fatal(http.ListenAndServe(":9999", r))
}
