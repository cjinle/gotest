package mylib

import (
	"fmt"
)

type UserProperty struct {
	Money int64
	Coin  int64
	Score int64
}

type UserMap2 map[int]*UserProperty

func MoneyServer2() {
	u := make(UserMap2)
	u[123] = &UserProperty{100, 0, 0}

	ul := []UserMap2{u}

	fmt.Println(*u[123])
	fmt.Println(ul[0][123].Money)

	u[123].Money = 1000

	fmt.Println(*u[123])
	fmt.Println(ul[0][123].Money)
}
