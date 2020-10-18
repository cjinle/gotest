package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// create an random array
func RandArray(n int) []int {
	rand.Seed(time.Now().UnixNano())

	arr := make([]int, n)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(n)
	}

	return arr
}

func Assert(t bool, s string) {
	if !t {
		panic(s)
	}
}

func Implode(sep string, v ...interface{}) string {
	data := make([]string, len(v))
	for idx, val := range v {
		data[idx] = fmt.Sprint(val)
	}
	return strings.Join(data, sep)
}

func Implode2(sep string, arr []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(arr)), "&"), "[]")
}

// Md5Sum 给字符串加密md5
func Md5Sum(s string) string {
	bytes := md5.Sum([]byte(s))
	// return fmt.Sprintf("%x", bytes)
	return hex.EncodeToString(bytes[:])
}

func Hello() string {
	return "hello world!"
}
