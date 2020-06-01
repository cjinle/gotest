package main

import (
	"fmt"

	mylib "github.com/cjinle/test/mylib"
)

func main() {
	fmt.Println("test main starting ... ")
	mylib.ArrOuput()
	mylib.Http()
	mylib.LogOutput()
	mylib.Redis()
}
