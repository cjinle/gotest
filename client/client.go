package client

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
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

	buff = bytes.NewBuffer(data)

	v := DataUnpack(bytes.NewBuffer(rb[16:recvLen]))
	fmt.Println(v)

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


func DataPack(buff *bytes.Buffer) []byte {
	ret := bytes.NewBuffer([]byte{})
	binary.Write(ret, binary.BigEndian, uint32(buff.Len()+16))
	binary.Write(ret, binary.BigEndian, byte(0x20))
	binary.Write(ret, binary.BigEndian, byte(0x21))
	binary.Write(ret, binary.BigEndian, uint16(1))
	binary.Write(ret, binary.BigEndian, byte(0x10))
	random := rand.Intn(1000)
	binary.Write(ret, binary.BigEndian, uint32(0x1004^random))
	binary.Write(ret, binary.BigEndian, uint16(random))
	binary.Write(ret, binary.BigEndian, byte(0))
	binary.Write(ret, binary.BigEndian, buff.Bytes())
	return ret.Bytes()
}