package main

import (
	"log"
)

// Param 参数
type Param struct {
	API     int    `json:"api"`
	Token   string `json:"token"`
	Content string `json:"content"`
}

func main() {
	log.Println("push json-rpc")
	// ioutil.ReadAll()
}
