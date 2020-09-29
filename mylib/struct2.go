package mylib

import (
	"fmt"
)

type S struct {
	int
	string
}

func Struct2() {
	s := S{1, "hello"}
	fmt.Println(s)
	fmt.Println(s.int)
	fmt.Println(s.string)
}
