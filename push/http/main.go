package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Param 参数
type Param struct {
	API     int    `json:"api"`
	Token   string `json:"token"`
	Content string `json:"content"`
}

// Result 返回值
type Result struct {
	Ret int `json:"ret"`
}

// DefaultHandle 默认处理函数
func DefaultHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("please use /push"))
}

var err error

func main() {
	log.Println("push http server start ... ")
	http.HandleFunc("/", DefaultHandle)
	http.HandleFunc("/push", func(w http.ResponseWriter, r *http.Request) {
		// log.Println(r.URL, r.FormValue("foo"), r.UserAgent())
		// err = r.ParseForm()
		// if err != nil {
		// 	log.Println(err)
		// }
		// log.Println("--->", r.FormValue("foo"), r.PostForm)
		enc := json.NewEncoder(w)
		// res := &Result{Ret: 0}
		// enc.Encode(res)
		now := int(time.Now().Unix())
		log.Println(now)
		enc.Encode(&Result{Ret: now})
	})
	http.ListenAndServe("0.0.0.0:6060", nil)
}
