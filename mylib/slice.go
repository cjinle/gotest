package mylib

import (
	"fmt"
)

func Slice() {
	s := []string{"a", "b", "c"}
	fmt.Println(s)
	s = s[0:0]
	fmt.Println(s)
}
