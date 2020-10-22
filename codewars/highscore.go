package codewars

import (
	"fmt"
	"sort"
	"strings"
)

func HighScore(s string) string {
	c := make(map[byte]int)

	for _, v := range []byte(s) {
		if _, ok := c[v]; ok {
			c[v]++
		} else {
			c[v] = 1
		}
	}
	out := ""
	for {
		if len(c) == 0 {
			break
		}
		maxNum := 0
		maxIdx := byte(0)
		for idx, num := range c {
			if maxNum < num {
				maxNum = num
				maxIdx = idx
			}
		}
		delete(c, maxIdx)
		out = out + fmt.Sprintf("%c=%d, ", maxIdx, maxNum)
	}
	out = strings.Trim(out, ", ")
	return out
}

func HighScore2(s string) string {
	arr := strings.Split(s, " ")
	maxIdx, maxScore := 0, 0
	for k, v := range arr {
		score := 0
		for _, b := range []byte(v) {
			score += int(b) - 96
		}
		if maxScore < score {
			maxScore = score
			maxIdx = k
		}
	}
	return arr[maxIdx]
}

type Item struct {
	K byte
	V int
}

func HighScore3(s string) string {
	c := make(map[byte]int)

	for _, v := range []byte(s) {
		if _, ok := c[v]; ok {
			c[v]++
		} else {
			c[v] = 1
		}
	}
	data := []Item{}
	for k, v := range c {
		data = append(data, Item{k, v})
	}
	out := ""
	sort.SliceStable(data, func(i, j int) bool { return data[i].V > data[j].V })
	fmt.Println(data)

	for _, v := range data {
		out = out + fmt.Sprintf("%c=%d, ", v.K, v.V)
	}
	out = strings.Trim(out, ", ")
	return out
}
