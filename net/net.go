package net


import (
	"log"
	"net/http"
	"io/ioutil"
)

func HttpGet() {
	url := "http://v.baidu.com/channel/short/newxiaopin?format=json&pn=1&_=1607419629"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := ioutil.ReadAll(res.Body)

	log.Println(string(bytes))
}