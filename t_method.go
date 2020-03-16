package main

import "fmt"

type MyInt int

func (i MyInt) Test() int {
	fmt.Printf("Type: %T Value: %v\n", i, i)
	return int(i) + 10
}

func main() {
	aa := MyInt(123)

	fmt.Println(aa.Test())
}
