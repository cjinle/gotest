package main

import "fmt"

func main() {
	x := []int{1, 2, 3}
	y := []int{4, 5, 6}
	x = append(x, y...)

	for i := 7; i < 20; i++ {
		x = append(x, i)
	}
	fmt.Println(x)
}
