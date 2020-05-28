package mylib

import (
	"encoding/binary"
	"fmt"
	"net"
)

func Leaf() {
	conn, err := net.Dial("tcp", "127.0.0.1:3563")
	if err != nil {
		panic(err)
	}

	data := []byte(`{
		"Hello": {
			"Name": "chenjinle ... ... "
		}
	}`)

	m := make([]byte, 2+len(data))
	fmt.Println(data)
	fmt.Println(m)
	binary.BigEndian.PutUint16(m, uint16(len(data)))

	copy(m[2:], data)
	fmt.Println(m)

	conn.Write(m)

}
