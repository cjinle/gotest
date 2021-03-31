package main

import (
	"log"
	"time"
)

func main() {
	chan1 := make(chan struct{})
	chan2 := make(chan struct{})

	// chan1 <- struct{}{}
	example1(chan1, chan2)
	// example2(chan1, chan2)

	for {
		chan1 <- struct{}{}
		time.Sleep(time.Second)

	}
}

func example1(chan1, chan2 chan struct{}) {
	go func() {
		for {
			select {
			case x := <-chan1:
				log.Println("chan1 ... ")
				chan2 <- x
				log.Println("chan1 ... done")
			}
		}
	}()

	go func() {
		for {
			select {
			case <-chan2:
				log.Println("chan2 ...")
			}
		}
	}()
}

func example2(chan1, chan2 chan struct{}) {
	go func() {
		for {
			select {
			case x := <-chan1:
				log.Println("chan1 ... ")
				chan2 <- x
				log.Println("chan1 ... done")
			case <-chan2:
				log.Println("chan2 ...")
			}
		}
	}()

}
