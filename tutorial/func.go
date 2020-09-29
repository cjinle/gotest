package tutorial

import (
	"fmt"
	// "strings"
)

type ActInfo []string

func Func() {
	var s = make(map[rune]string)
	// s = strings.Map(func(r rune) rune {
	// 	return r + 1
	// }, "HAL-9000")
	s[4] = "bar"
	fmt.Println(s)
	var a rune
	a = 4
	fmt.Println(s[a])
	// fmt.Printf("%v", s(a))
	fmt.Println(Test())
}

func Test() ActInfo {
	return ActInfo{"aaaa", "bbb"}
}
