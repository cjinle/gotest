package codewars

import "fmt"

func FindOutlier(integers []int) int {
	arr := [4]int{0, 0, 0, 0}
	fmt.Println(integers)
	for _, v := range integers {
		if v < 0 {
			v *= -1
		}
		n := v % 2
		arr[n]++
		arr[n+2] = v
	}
	fmt.Println(arr)
	if arr[0] == 1 {
		return arr[2]
	}
	return arr[3]
}
