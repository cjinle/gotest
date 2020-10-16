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


// ---------- method -------------
type MyInt int

func (i MyInt) Test() int {
	fmt.Printf("Type: %T Value: %v\n", i, i)
	return int(i) + 10
}

func Method() {
	aa := MyInt(123)

	fmt.Println(aa.Test())
}


// ---------------------------------
type P int

func (P) f() {
	fmt.Println("f call!")
}

type Point struct {
	X, Y float64
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func Method2() {
	var x P
	x.f()

	var cp ColoredPoint

	cp.X = 1.1
	cp.Y = 1.2

	fmt.Println(cp)

}
