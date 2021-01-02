package client

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net"
	"time"
)

type Client struct {
	conn *net.TCPConn
}

func NewClient() (*Client, error) {
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", "192.168.100.107:8051")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		panic(err)
	}
	return &Client{conn}, err
}

func (client *Client) Close() {
	client.conn.Close()
}

func (client *Client) GetGameInfo(mid uint32) interface{} {
	// c, _ := net.Dial("tcp", "192.168.100.107:8051")
	// defer c.Close()

	conn := client.conn
	buff := bytes.NewBuffer([]byte{})
	binary.Write(buff, binary.BigEndian, uint32(mid))
	dp := DataPack(buff, 0x1004)
	fmt.Printf("%s", hex.Dump(dp))

	n, err := conn.Write(dp)
	fmt.Println(n, err)

	headData := make([]byte, 24)
	_, err = io.ReadFull(conn, headData)
	fmt.Printf("%s", hex.Dump(headData))

	var ret, dataLen int32
	buff = bytes.NewBuffer(headData[16:])
	binary.Read(buff, binary.BigEndian, &ret)
	binary.Read(buff, binary.BigEndian, &dataLen)
	fmt.Println(ret, dataLen)

	data := make([]byte, dataLen)
	_, err = io.ReadFull(conn, data)
	if err != nil {
		fmt.Println(err)
	}
	buff = bytes.NewBuffer(data)
	data = make([]byte, dataLen-1)

	binary.Read(buff, binary.BigEndian, &data)
	fmt.Printf("%s", hex.Dump(data))

	var v interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		fmt.Println(err)
	}
	return v
}

func (client *Client) UpdateMoney(mid, mode, num int) interface{} {
	conn := client.conn
	buff := bytes.NewBuffer([]byte{})
	binary.Write(buff, binary.BigEndian, uint32(mid))
	strArr := []byte{'1'}
	binary.Write(buff, binary.BigEndian, uint32(len(strArr)+1))
	binary.Write(buff, binary.BigEndian, strArr)
	binary.Write(buff, binary.BigEndian, uint8(0))
	binary.Write(buff, binary.BigEndian, int8(1))
	binary.Write(buff, binary.BigEndian, uint32(1))
	binary.Write(buff, binary.BigEndian, uint64(num))
	binary.Write(buff, binary.BigEndian, uint32(mode))

	dp := DataPack(buff, 0x1005)
	fmt.Printf("%s", hex.Dump(dp))

	n, err := conn.Write(dp)
	fmt.Println(n, err)

	headData := make([]byte, 24)
	_, err = io.ReadFull(conn, headData)
	fmt.Printf("%s", hex.Dump(headData))

	var ret, dataLen int32
	buff = bytes.NewBuffer(headData[16:])
	binary.Read(buff, binary.BigEndian, &ret)
	binary.Read(buff, binary.BigEndian, &dataLen)
	// fmt.Println(ret, dataLen)

	data := make([]byte, dataLen)
	_, err = io.ReadFull(conn, data)
	if err != nil {
		fmt.Println(err)
	}
	buff = bytes.NewBuffer(data)
	data = make([]byte, dataLen-1)

	binary.Read(buff, binary.BigEndian, &data)
	fmt.Printf("%s", hex.Dump(data))

	var v interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		fmt.Println(err)
	}
	return v
}

func DataUnpack(buff *bytes.Buffer) interface{} {
	var ret, dataLen int32
	binary.Read(buff, binary.BigEndian, &ret)
	binary.Read(buff, binary.BigEndian, &dataLen)
	info := make([]byte, dataLen-1)
	binary.Read(buff, binary.BigEndian, &info)
	var v interface{}
	err := json.Unmarshal(info, &v)
	if err != nil {
		fmt.Println(err)
	}
	return v
}

func DataPack(buff *bytes.Buffer, cmd int) []byte {
	ret := bytes.NewBuffer([]byte{})
	binary.Write(ret, binary.BigEndian, uint32(buff.Len()+16))
	binary.Write(ret, binary.BigEndian, byte(0x20))
	binary.Write(ret, binary.BigEndian, byte(0x21))
	binary.Write(ret, binary.BigEndian, uint16(1))
	binary.Write(ret, binary.BigEndian, byte(0x10))
	rand.Seed(time.Now().Unix())
	random := rand.Intn(1000)
	random = 100
	binary.Write(ret, binary.BigEndian, uint32(cmd^random))
	binary.Write(ret, binary.BigEndian, uint16(random))
	binary.Write(ret, binary.BigEndian, byte(0))
	binary.Write(ret, binary.BigEndian, buff.Bytes())
	return ret.Bytes()
}
