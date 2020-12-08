package main

import (
	"encoding/binary"
	"fmt"
)


func main() {
	fmt.Println("test main starting ... ")
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, 1)
	fmt.Println(bs)
	binary.BigEndian.PutUint32(bs, 1)
	fmt.Println(bs)


	var aa int
	aa = 0x1001
	bb := int(017)
	fmt.Println(aa, aa == 4097, bb)
}

