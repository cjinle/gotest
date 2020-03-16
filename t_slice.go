package main

import (
	"fmt"
)

func main() {
	s := []string{"a", "b", "c"}
	fmt.Println(s)
	s = s[0:0]
	fmt.Println(s)
}
