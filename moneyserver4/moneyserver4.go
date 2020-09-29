package moneyserver4

import (
	"fmt"
	"sync"
)

type User struct {
	uid   uint32
	money uint64
}

type UserMap struct {
	mu    sync.Mutex
	users map[uint32]*User
}

const NUM = 100

var userMaps [NUM]*UserMap

func MoneyServer4() {
	for i := range userMaps {
		userMaps[i] = &UserMap{users: make(map[uint32]*User, 500)}
	}
	fmt.Println(userMaps)
}

func get(uid uint32) *User {
	userMap := userMaps[uid%NUM]
	userMap.mu.Lock()
	user := userMap.users[uid]
	userMap.mu.Unlock()
	return user
}
