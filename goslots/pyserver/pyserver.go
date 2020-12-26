package main

import (
	"fmt"
	"net"
	"strconv"

	"github.com/cjinle/test/goslots"
)

type User struct {
	Id    int
	Money int
	Win   int
	Lose  int
}

var ConnMap map[string]net.Conn

func main() {
	fmt.Println("slots server start ... ")
	ConnMap = make(map[string]net.Conn)
	netListen, err := net.Listen("tcp", ":1234")
	CheckError(err)
	defer netListen.Close()
	startId := 1

	fmt.Println("Waiting for clients")
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}

		user := &User{startId, 500000, 0, 0}
		startId++

		fmt.Println(conn.RemoteAddr().String(), " tcp connect success")
		ConnMap[conn.RemoteAddr().String()] = conn
		go handleConnection(conn, user)
	}
}

func handleConnection(conn net.Conn, user *User) {
	conn.Write([]byte("Welcome: " + UserString(user)))
	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			return
		}
		bet, _ := strconv.Atoi(string(buffer[:n]))
		fmt.Printf("ID: %d, Bet: %d\n", user.Id, bet)
		if bet > user.Money {
			fmt.Printf("ID: %d, not enough money.\n", user.Id, bet)
			fmt.Println(UserString(user))
			return
		}
		user.Money -= bet
		win, value := goslots.Bet(bet)
		fmt.Println(win, value)
		if win > 0 {
			user.Money += win
			user.Win++
		} else {
			user.Lose++
		}
		conn.Write([]byte(UserString(user)))
	}
}

func UserString(user *User) string {
	return fmt.Sprintf("ID: %d, Money: %d, Win: %d, Lose: %d", user.Id, user.Money, user.Win, user.Lose)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
