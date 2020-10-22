package main

import (
	"crypto/rand"
	"fmt"
	"strings"

	"github.com/cjinle/test/utils"
)

// Implode str...
func Implode(sep string, v ...interface{}) string {
	data := make([]string, len(v))
	for idx, val := range v {
		data[idx] = fmt.Sprint(val)
	}
	return strings.Join(data, sep)
}

// Explode sth...
func Explode(seq, str string) []string {
	return []string{"hello", "word"}
}

func main() {
	fmt.Println(utils.Md5Sum("123456"))

	A := []int{1, 2, 3}
	fmt.Println(utils.Implode2("&", A))

	bytes := make([]byte, 16)
	rand.Read(bytes)
	fmt.Printf("%x", bytes)
	utils.Assert(1 == 1, "1 == 1")
	utils.Assert(false == false, "false is false")

	
	s.replace(/\|$/, '').split('|')
}
