package csv

import (
	"log"
	"testing"
)

func TestExample(t *testing.T) {
	log.Println("test start ... ")
	// ExampleReader()
	ExampleWriter()
}

func TestOne(t *testing.T) {
	log.Println("test one start ... ")

	var v interface{}
	v = float64(999)
	log.Println(v)
	switch v.(type) {
	case float64:
		log.Println(v)
	}
	// log.Println(v.(type))
}
