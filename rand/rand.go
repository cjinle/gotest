package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	Func1()
}

func Func1() {
	rand.Seed(time.Now().Unix())
	arr := rand.Perm(52)
	fmt.Println(arr)
}
