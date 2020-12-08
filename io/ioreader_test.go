package io

import (
	"log"
	"testing"
	"io/ioutil"
)

func TestRead(t *testing.T) {
	mydata := &MyData{}
	// log.Println(mydata.Read([]byte("hello")))
	log.Println(ioutil.ReadAll(mydata))
}