package a

import (
	"fmt"
)

func init() {
	fmt.Println("b.go func init ... ")
}

func B() {
	fmt.Println("b.go func B ... ")
}
