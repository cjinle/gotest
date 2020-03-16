package main

import (
	"fmt"
	"time"
	// "github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/load"
)

func main() {

	// v, _ := mem.VirtualMemory()
	// fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
	// fmt.Println(v)
	for {
		v2, _ := load.Avg()
		fmt.Printf("%v,%v,%v\n", v2.Load1, v2.Load5, v2.Load15)
		fmt.Println(v2)
		time.Sleep(5 * time.Second)
	}
	// v2, _ := load.Avg()
	// fmt.Println(v2)
}
