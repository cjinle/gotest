package io

import (
	"log"
	"errors"
)

type MyData struct {

}

func (md *MyData) Read(p []byte) (int, error) {
	log.Println(string(p))

	return 0, errors.New("EOF")
}

