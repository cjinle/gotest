package io

import (
	"log"
	// "errors"
	"io"
)

type MyData struct {
	r  *io.Reader
}

func (md *MyData) Read(p []byte) (int, error) {
	log.Println(string(p))

	return 0, nil
	// return 0, errors.New("EOF")
}

