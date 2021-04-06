package proc

import (
	"fmt"
	"runtime"
	"time"
)

func MultiProc() {
	runtime.GOMAXPROCS(2)
	go func() {
		for {
			fmt.Println("go1 ...")
			time.Sleep(time.Second)
		}
	}()
	go func() {
		for {
			fmt.Println("go2 ...")
			time.Sleep(time.Second)
		}
	}()
	fmt.Println("after go")
	select {}
}

func SingleProc() {
	runtime.GOMAXPROCS(1)
	go func() {
		for {
			fmt.Print(".")
			time.Sleep(5000 * time.Microsecond)
		}
	}()
	time.Sleep(time.Second)
	fmt.Println("done")
	select {}
}
