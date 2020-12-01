package tutorial

import (
	"fmt"
	"testing"
)

const TimeOut int64 = 200

func TestOne(t *testing.T) {
	fmt.Println("test start ...")

	// DownloadPic()
	// UdpLogs()
	// Tick()
	// OpTime()
	// fmt.Printf("%T %v", TimeOut, TimeOut)
	// fmt.Println(GetTimeStr())
	MyJson()

	Struct2()
}

// func TestBuffer(t *testing.T) {
// 	Buffer()
// }

// func TestArr(t *testing.T) {
// 	ArrOuput()
// }

// func TestTime(t *testing.T) {
// 	Time()
// }

// func TestMysql(t *testing.T) {
// 	Mysql()
// }
