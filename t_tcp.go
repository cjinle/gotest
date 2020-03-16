package main

import (
	"fmt"
	"net"
	"os"
)

var ConnMap map[string]net.Conn

func main() {
	ConnMap = make(map[string]net.Conn)
	netListen, err := net.Listen("tcp", ":1234")
	CheckError(err)

	defer netListen.Close()

	Log("Waiting for clients")
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}

		Log(conn.RemoteAddr().String(), " tcp connect success")
		ConnMap[conn.RemoteAddr().String()] = conn
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			Log(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}
		Log(conn.RemoteAddr().String(), "receive data length:", n)
		Log(conn.RemoteAddr().String(), "receive data:", buffer[:n])
		Log(conn.RemoteAddr().String(), "receive data string:", string(buffer[:n]))

		n, err = conn.Write([]byte("server response\n"))
		if err != nil {
			fmt.Println("write error:", err)
		} else {
			fmt.Println("write bytes, content")
		}
		SendToClient(string(buffer[:n]))
	}
}

func Log(v ...interface{}) {
	fmt.Println(v...)
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func SendToClient(msg string) {
	b := []byte(msg)
	for _, conn := range ConnMap {
		conn.Write(b)
	}
}
