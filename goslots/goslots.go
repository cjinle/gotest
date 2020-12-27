package goslots

import (
	// "fmt"
	"math/rand"
	"time"
)

const (
	MAX_VALUE_NUM = 9
	MAX_COL_NUM = 3
)

type Item struct {
	Times int
	Value []int
	Pr    float64
}

var Rule = []Item{
	Item{1000, []int{1, 1, 1}, 0.000001},
	Item{100, []int{2, 2, 2}, 0.00005},
	Item{50, []int{3, 3, 3}, 0.0001},
	Item{50, []int{4, 4, 4}, 0.0001},
	Item{30, []int{5, 5, 5}, 0.0002},
	Item{20, []int{6, 6, 6}, 0.0002},
	Item{10, []int{7, 7, 7}, 0.0005},
	Item{8, []int{8, 8, 8}, 0.005},
	Item{5, []int{9, 9, 9}, 0.01},
	Item{3, []int{1, 1}, 0.015},
	Item{3, []int{2, 2}, 0.015},
	Item{3, []int{3, 3}, 0.015},
	Item{3, []int{4, 4}, 0.015},
	Item{2, []int{5, 5}, 0.04},
	Item{2, []int{6, 6}, 0.04},
	Item{2, []int{7, 7}, 0.04},
	Item{2, []int{8, 8}, 0.04},
	Item{2, []int{9, 9}, 0.04},
}

func Lucky() Item {
	minPr, multiple := 1.0, 1.0
	for _, v := range Rule {
		if v.Pr <= 0 {
			continue
		}
		if minPr > v.Pr {
			minPr = v.Pr
		}
	}
	for minPr < 1 {
		minPr *= 10
		multiple *= 10
	}

	rand.Seed(time.Now().UnixNano())
	randNum := float64(rand.Intn(int(multiple)))

	data := Item{}
	var start float64 = 0
	for _, v := range Rule {
		start += v.Pr * multiple
		if randNum <= start {
			data = v
			break
		}
	}

	fixNum := MAX_COL_NUM - len(data.Value)
	for fixNum > 0 {
		randValue := rand.Intn(MAX_VALUE_NUM) + 1
		if !inArray(randValue, data.Value) {
			data.Value = append(data.Value, randValue)
			fixNum--
		}
	}

	rand.Shuffle(len(data.Value), func(i, j int) { data.Value[i], data.Value[j] = data.Value[j], data.Value[i] })
	return data
}

func inArray(v int, arr []int) bool {
	for _, val := range arr {
		if val == v {
			return true
		}
	}
	return false
}

func Bet(money int) (int, []int) {
	win := 0
	ret := []int{}
	if money <= 0 {
		return win, ret
	}
	item := Lucky()
	// fmt.Println(item)
	ret = item.Value
	if item.Times > 0 {
		win = money * item.Times
	}
	return win, ret
}
