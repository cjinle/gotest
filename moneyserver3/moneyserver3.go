package moneyserver3

import (
	"fmt"
	"math/rand"
	"sync"
)

const TABLENUM int = 100

type UserInfo struct {
	Money int
}

type UserList [TABLENUM]map[int]*UserInfo

var (
	ul UserList
	mu sync.Mutex
)

func MoneyServer3() {
	ul.InitData()
	for i := 0; i < 99999; i++ {
		ul.AddMoney(i%100, rand.Intn(999999))
	}
	fmt.Println(ul)
}

func (ul *UserList) InitData() {
	for i := 0; i < TABLENUM; i++ {
		ul[i] = map[int]*UserInfo{}
	}
	return
}

func (ul *UserList) GetTable(uid int) int {
	idx := uid % TABLENUM
	if _, ok := ul[idx][uid]; !ok {
		ul[idx][uid] = &UserInfo{0}
	}
	return idx
}

func (ul *UserList) AddMoney(uid int, num int) (int, error) {
	idx := ul.GetTable(uid)
	mu.Lock()
	ul[idx][uid].Money = ul[idx][uid].Money + num
	mu.Unlock()
	return ul[idx][uid].Money, nil
}

func (ul *UserList) GetMoney(uid int) (int, error) {
	idx := ul.GetTable(uid)
	return ul[idx][uid].Money, nil
}
