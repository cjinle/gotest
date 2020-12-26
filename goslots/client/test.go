package main

import (
	"fmt"
	"net"
	"time"

	// "github.com/cjinle/goslots"
	"github.com/cjinle/test/goslots/pb"
	"github.com/golang/protobuf/proto"
)

func main() {
	fmt.Println("slots client start ... ")
	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	CheckError(err)

	var buf [1024]byte
	n, err := conn.Read(buf[0:])
	CheckError(err)
	fmt.Println(n, string(buf[:n]))

	bet := &pb.Bet{Money: 50000}
	bytes, err := proto.Marshal(bet)
	CheckError(err)
	for {
		conn.Write([]byte(bytes))
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		result := &pb.Result{}
		err = proto.Unmarshal(buf[:n], result)
		if err != nil {
			return
		}
		fmt.Println(result.String())

		time.Sleep(time.Second)
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
