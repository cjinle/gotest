package encode

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/rand"
)

func BinPacks() {
	buff := bytes.NewBuffer([]byte{})
	binary.Write(buff, binary.BigEndian, uint32(20))
	binary.Write(buff, binary.BigEndian, byte(0x20))
	binary.Write(buff, binary.BigEndian, byte(0x21))
	binary.Write(buff, binary.BigEndian, uint16(1))
	binary.Write(buff, binary.BigEndian, byte(0x10))
	random := rand.Intn(1000)
	binary.Write(buff, binary.BigEndian, uint32(0x1004^random))
	binary.Write(buff, binary.BigEndian, byte(0))
	binary.Write(buff, binary.BigEndian, uint32(2494))
	fmt.Printf("%s", hex.Dump(buff.Bytes()))
	fmt.Println(buff.Bytes(), random, binary.Size(byte(0x10)))

}
