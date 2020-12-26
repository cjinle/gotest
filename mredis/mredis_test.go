package mredis

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestConn(t *testing.T) {
	fmt.Println(".....")

	// a := 23123
	a := "123.12"
	var b interface{} = a
	v := reflect.ValueOf(b)
	fmt.Println(v.Type())

	r := NewRedis("127.0.0.1:6379")
	fmt.Println(r.Get("foo", "string"))
	fmt.Println(r.Get("foo2", "string"))
	fmt.Println(r.Get("foo3", "int64"))
	fmt.Println(r.Set("foo3", "123123123"))
	result, err := r.Get("foo4", "int64")
	if err != nil {
		// panic(err)
		fmt.Println(err)
	}
	// result.(int) += 100
	fmt.Printf("result type = %T, value = %v \n", result, result)
	fmt.Println(r.Set("foo3", 888))
	// r.Close()
	for {
		select {
		case <-time.After(time.Second * 1):
			fmt.Println(r.Ping())
			// return
		}
	}

}
