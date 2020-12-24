package csv

import (
	"fmt"
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
	var xx byte
	xx = byte(12)
	xx += byte(1)
	log.Printf("%T, %v", xx, xx)

	yy := uint32(99)
	log.Println(yy, fmt.Sprint(yy))
}
