package mylib

import (
	"fmt"
	"math/rand"
)

const TABLENUM int = 100

type MoneyServer interface {
	InitData()
	GetTable(uid int) int
	AddMoney(uid int, num int) (int, error)
	GetMoney(uid int) (int, error)
	String() string
}

type UserInfo struct {
	Money int
}

type UserList [TABLENUM]map[int]*UserInfo

func MoneyServer() {
	var ms MoneyServer = &UserList{}
	ms.InitData()
	// test
	for i := 0; i < 999; i++ {
		ms.AddMoney(i%100, rand.Intn(999999))
	}

	fmt.Println(ms)
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
	ul[idx][uid].Money = ul[idx][uid].Money + num
	return ul[idx][uid].Money, nil
}

func (ul *UserList) GetMoney(uid int) (int, error) {
	idx := ul.GetTable(uid)
	return ul[idx][uid].Money, nil
}

func (ul *UserList) String() string {
	for i := 0; i < TABLENUM; i++ {
		if len(ul[i]) == 0 {
			continue
		}
		for k, v := range ul[i] {
			fmt.Printf("%d\t%v\n", k, v.Money)
		}
	}
	return "done"
}
