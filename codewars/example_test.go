package codewars

import (
	"fmt"
	"testing"
	"sort"
)

func _TestTwoToOne(t *testing.T) {
	var a, b string
	a = "xyaabbbccccdefww"
	b = "xxxxyyyyabklmopq"
	fmt.Println(TwoToOne2(a, b))


	arr := []int{1,2,3,98,10,5,18}
	sort.Ints(arr)
	fmt.Println(arr)
	
	fmt.Println(sort.Reverse(sort.IntSlice(arr)))

}

func TestSth(t *testing.T) {
	n := 1907
	fmt.Println(int(n/100))
}