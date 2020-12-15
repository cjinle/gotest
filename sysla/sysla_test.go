package sysla_test

import (
	"log"
	"time"

	// "github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/load"
)

func TestAvg() {

	// v, _ := mem.VirtualMemory()
	// fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
	// fmt.Println(v)
	for {
		v2, _ := load.Avg()
		log.Printf("%v,%v,%v\n", v2.Load1, v2.Load5, v2.Load15)
		log.Println(v2)
		time.Sleep(5 * time.Second)
	}
	// v2, _ := load.Avg()
	// fmt.Println(v2)
}
