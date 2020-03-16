package main

import (
	"fmt"
	// "io/ioutil"
	"net"
	"os"
)

type udpParam struct {
	conn *net.UDPConn
	f    *os.File
}

func main() {
	f, _ := os.OpenFile("t_udplogs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer f.Close()

	fmt.Println("udp logs server start...")
	udpAddr, err := net.ResolveUDPAddr("udp4", ":13333")
	if err != nil {
		panic(err)
	}

	listen, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		panic(err)
	}
	defer listen.Close()
	up := udpParam{listen, f}

	for {
		handleUDPConection(&up)
	}
}

func handleUDPConection(up *udpParam) {
	buf := make([]byte, 1024)
	n, addr, err := up.conn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("UDP client: ", addr)
	fmt.Println("Get messages: ", string(buf[:n]))

	msg := append([]byte("IP: "), buf[:n]...)
	_, err = up.conn.WriteToUDP(msg, addr)
	if err != nil {
		fmt.Println(err)
	}
	go up.WriteLog(string(buf[:n]))
}

func (up *udpParam) WriteLog(s string) {
	up.f.WriteString(s + "\n")
	return
}
