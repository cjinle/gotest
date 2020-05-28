package mylib

import (
	"fmt"
	"image/color"
)

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
