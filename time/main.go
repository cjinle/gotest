package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	var i int
	go func() {
		for {
			i++
			ch <- i
		}
	}()
	// test1(ch)
	test2(ch)
}

func test1(ch chan int) {
	after := time.NewTimer(5 * time.Second)
	for {
		select {
		case <-ch:
			fmt.Print(".")
			time.Sleep(time.Second)
		case <-after.C:
			fmt.Println(time.Now().Unix())
			break
		}
	}
}

func test2(ch chan int) {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ch:
			fmt.Print(".")
			time.Sleep(50000 * time.Microsecond)
		case <-ticker.C:
			fmt.Println(time.Now().Unix())
		}
	}
}
