package main


import (
	"fmt"
)

type S struct {
	int
	string
}

func main() {
	s := S{1,"hello"}
	fmt.Println(s)
	fmt.Println(s.int)
	fmt.Println(s.string)
}