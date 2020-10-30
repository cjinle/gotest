package client

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"encoding/hex"
	"net"
)

func GetGameInfo(mid uint32) {
	c, _ := net.Dial("tcp", "192.168.100.107:8051")
	defer c.Close()

	buff := bytes.NewBuffer([]byte{})
	binary.Write(buff, binary.BigEndian, uint32(mid))
	dp := DataPack(buff)
	fmt.Printf("%s", hex.Dump(dp))
	
	n, err := c.Write(dp)
	fmt.Println(n, err)

	rb := make([]byte, 4096)
	recvLen, err := c.Read(rb)
	if err != nil {
		fmt.Println(err)
	}
	
	data := rb[16:recvLen]
	fmt.Printf("%s", hex.Dump(data))


}

func GetGameInfo2(mid uint32) {
	c, _ := net.Dial("tcp", "192.168.100.107:8051")
	defer c.Close()

	buff := bytes.NewBuffer([]byte{})
	binary.Write(buff, binary.BigEndian, uint32(20))
	binary.Write(buff, binary.BigEndian, byte(0x20))
	binary.Write(buff, binary.BigEndian, byte(0x21))
	binary.Write(buff, binary.BigEndian, uint16(1))
	binary.Write(buff, binary.BigEndian, byte(0x10))
	// random := rand.Intn(1000)
	random := 100
	binary.Write(buff, binary.BigEndian, uint32(0x1004^random))
	binary.Write(buff, binary.BigEndian, uint16(random))
	binary.Write(buff, binary.BigEndian, byte(0))
	binary.Write(buff, binary.BigEndian, uint32(mid))

	fmt.Printf("%s", hex.Dump(buff.Bytes()))
	
	n, err := c.Write(buff.Bytes())
	fmt.Println(n, err)

	rb := make([]byte, 4096)
	// rb := []byte{}
	if _, err := c.Read(rb); err != nil {
		fmt.Println(err)
	}
	fmt.Println(rb)
}


func DataPack(buff *bytes.Buffer) []byte {
	retBuff := bytes.NewBuffer([]byte{})
	binary.Write(retBuff, binary.BigEndian, uint32(buff.Len()+16))
	binary.Write(retBuff, binary.BigEndian, byte(0x20))
	binary.Write(retBuff, binary.BigEndian, byte(0x21))
	binary.Write(retBuff, binary.BigEndian, uint16(1))
	binary.Write(retBuff, binary.BigEndian, byte(0x10))
	random := rand.Intn(1000)
	binary.Write(retBuff, binary.BigEndian, uint32(0x1004^random))
	binary.Write(retBuff, binary.BigEndian, uint16(random))
	binary.Write(retBuff, binary.BigEndian, byte(0))
	binary.Write(retBuff, binary.BigEndian, buff.Bytes())
	return retBuff.Bytes()
}