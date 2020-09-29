package tutorial

import (
	"fmt"
)

type UserMap struct {
	uid   int
	money int
}

func Make() {
	usermap := make([]int, 10)
	fmt.Println(usermap)

}
