package main

import (
	"fmt"
	"os"
)

func main() {
	var file string
	file = "aa"

	fi, err := os.Stat(file)

	fmt.Println(fi.Name(), fi.IsDir(), fi.Mode().String(), err)
}
