package main

import (
	"fmt"
	"net"

	"github.com/cjinle/test/goslots"
	"github.com/cjinle/test/goslots/pb"
	"github.com/golang/protobuf/proto"
)

// User struct
type User struct {
	ID    int32
	Money int32
	Win   int32
	Lose  int32
}

// ConnMap map
var ConnMap map[string]net.Conn

func main() {
	fmt.Println("slots server start ... ")
	ConnMap = make(map[string]net.Conn)
	netListen, err := net.Listen("tcp", ":1234")
	CheckError(err)
	defer netListen.Close()
	var startID int32 = 1

	fmt.Println("Waiting for clients")
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}

		user := &pb.User{Id: startID, Money: 500000, Win: 0, Lose: 0}
		startID++

		fmt.Println(conn.RemoteAddr().String(), " tcp connect success")
		ConnMap[conn.RemoteAddr().String()] = conn
		go handleConnection(conn, user)
	}
}

func handleConnection(conn net.Conn, user *pb.User) {
	conn.Write([]byte("Welcome: " + user.String()))
	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			return
		}
		bet := &pb.Bet{}
		err = proto.Unmarshal(buffer[:n], bet)
		if err != nil {
			return
		}

		fmt.Printf("ID: %d, Bet: %d\n", user.Id, bet.Money)
		if bet.Money > user.Money {
			fmt.Printf("ID: %d, %d not enough money.\n", user.Id, bet.Money)
			fmt.Println(user.String())
			conn.Close()
			return
		}
		user.Money -= bet.Money
		win, value := goslots.Bet(int(bet.Money))
		fmt.Println(win, value)
		if win > 0 {
			user.Money += int32(win)
			user.Win++
		} else {
			user.Lose++
		}
		result := &pb.Result{
			Ret:   0,
			Win:   int32(win),
			Value: []int32{int32(value[0]), int32(value[1]), int32(value[2])},
			User:  user,
		}
		bytes, err := proto.Marshal(result)
		if err != nil {
			return
		}
		conn.Write([]byte(bytes))
	}
}

// CheckError func
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
