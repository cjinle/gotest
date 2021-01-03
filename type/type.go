package main

import (
	"fmt"
	"reflect"
)

func main() {

	// n := 123123
	// n := make(map[string]int)
	var n interface{}
	n = 9999
	v := reflect.ValueOf(n)
	fmt.Println(v.Type())
}
