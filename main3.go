package main

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"strings"
)

func Implode(sep string, v ...interface{}) string {
	data := make([]string, len(v))
	for idx, val := range v {
		data[idx] = fmt.Sprint(val)
	}
	return strings.Join(data, sep)
}

func main() {
	xx := md5.Sum([]byte("hello"))
	fmt.Println(fmt.Sprintf("%x", xx))

	fmt.Println(md5str("123456"))

	A := []int{1, 2, 3}
	B := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(A)), "&"), "[]")
	fmt.Println(B)
	fmt.Println(Implode("&", A))

	bytes := make([]byte, 16)
	rand.Read(bytes)
	fmt.Printf("%x", bytes)

}

func md5str(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
