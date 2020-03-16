package main

import (
	"time"
	"fmt"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 50
	fibN := fib(n)
	fmt.Printf("\rFib(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}