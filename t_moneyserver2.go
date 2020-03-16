package main


import (
	"fmt"
)

type UserProperty struct {
	Money int64
	Coin int64
	Score int64
}

type UserMap map[int]*UserProperty


func main() {
	u := make(UserMap)
	u[123] = &UserProperty{100,0,0}

	ul := []UserMap{u}

	fmt.Println(*u[123])
	fmt.Println(ul[0][123].Money)

	u[123].Money = 1000

	fmt.Println(*u[123])
	fmt.Println(ul[0][123].Money)
}