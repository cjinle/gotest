package mylib

import (
	"fmt"
	"strconv"
)

type Currency int

const (
	USD Currency = iota // 美元
	EUR                 // 欧元
	GBP                 // 英镑
	RMB                 // 人民币
)

func str() {
	fmt.Println(USD, EUR, GBP, RMB)

	s := "Hello, 世界"

	fmt.Println(s, len(s))

	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	b := []byte(s)
	fmt.Println(string(b))

	fmt.Println(string(strconv.AppendInt(b, 666, 16)))

	// fmt.Println(basename("a/b/c.go")) // "c"
	// fmt.Println(basename("c.d.go"))   // "c.d"
	// fmt.Println(basename("abc"))      // "abc"
}

func basename(name string) string {
	i := len(name) - 1
	for ; i > 0 && name[i] == '/'; i-- {
		name = name[:i]
	}
	for i--; i >= 0; i-- {
		if name[i] == '/' {
			name = name[i+1:]
			break
		}
	}

	return name
}
