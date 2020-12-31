package goslots

import (
	"fmt"
	"testing"
)


func TestBet(t *testing.T) {
	//fmt.Println(Lucky())
	stat := make(map[int]int)
	maxNum := 1000000
	i := 0
	for i < maxNum {
		item := Lucky()
		stat[item.Times]++
		i++
	}
	for k, v := range stat {
		fmt.Println(k, v, float32(v)/float32(maxNum))	
	}
	// fmt.Println(stat)
}

func BenchmarkLucky(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Lucky()
	}
}