package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	xx := md5.Sum([]byte("hello"))
	fmt.Println(fmt.Sprintf("%x", xx))
}
