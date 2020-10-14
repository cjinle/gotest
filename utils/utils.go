package utils

import (
	"math/rand"
	"time"
)

// create an random array
func RandArray(n int) []int {
	rand.Seed(time.Now().UnixNano())

	arr := make([]int, n)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(n)
	}

	return arr
}