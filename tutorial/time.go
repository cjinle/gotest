package tutorial

import (
	"fmt"
	"time"
)

func Tick() {
	for range time.Tick(time.Second * 5) {
		fmt.Println("tick...")
	}
}

func OpTime() {
	now := time.Now()
	fmt.Println(time.Now().String(), time.Now().Location())
	time.Sleep(time.Second * 1)
	fmt.Println(time.Since(now).Seconds())
	time.Sleep(time.Second * 1)
	fmt.Println(now.Unix(), time.Now().Unix())
}

func GetTimeStr() string {
	now := time.Now()
	// ret := fmt.Sprintf("%d-%d-%d %d:%d:%d", now.Year(), now.Month(), 
		// now.Day(), now.Hour(), now.Minute(), now.Second())
	// ret := now.Format("2020-01-01 00:00:00")
	// ret := now.Format(time.RFC3339)
	ret := now.Format("2006-01-02 15:04:05")
	return ret
}

func GetTimeStr2() string {
	return time.Now().Format("20060102")
}
