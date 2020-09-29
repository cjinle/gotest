package tutorial

import "fmt"

func ArrOuput() {
	var a = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a)

	r := [...]int{99: -1}
	fmt.Println(r)

	fmt.Println("hello world")

	// var str = [...]byte{'a', 'b', 'c'}
	// str = append(str[:], 'x')
	// fmt.Println(str)
}
