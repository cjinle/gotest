package codewars

import (
	"fmt"
	"sort"
	"testing"
	// "strings"
)

func _TestTwoToOne(t *testing.T) {
	var a, b string
	a = "xyaabbbccccdefww"
	b = "xxxxyyyyabklmopq"
	fmt.Println(TwoToOne2(a, b))

	arr := []int{1, 2, 3, 98, 10, 5, 18}
	sort.Ints(arr)
	fmt.Println(arr)

	fmt.Println(sort.Reverse(sort.IntSlice(arr)))

}

func TestSth(t *testing.T) {
	// a := []int{121, 144, 19, 161, 19, 144, 19, 11}
	// b := []int{121, 14641, 20736, 361, 25921, 361, 20736, 361}
	// fmt.Println(Comp(a, b))

	// fmt.Println(HighScore2("take me to semynak"))
	fmt.Println(HighScore("bbaaabbccd"))
	fmt.Println(HighScore3("bbaaabbccd"))
	// fmt.Println(SplitStrings2("abcdef"))
	// fmt.Println(PickPeaks([]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3}))

	// x := []int{1,2,3}
	// for k, v := range x {
	// 	x[k] = v * 2
	// }
	// fmt.Println(x)
	// fmt.Println(WeirdCase2("Weird string case"), byte('A')+32)
	// fmt.Println(TwoSum([]int{1, 2, 3, 4}, 4))

	// a := byte('a') + byte(1)
	// b := rune('a') + rune(1)
	// fmt.Println(a, b)

	// var builder strings.Builder
	// builder.WriteString("hello world")
	// builder.WriteByte(97)
	// builder.WriteRune(rune('我'))
	// fmt.Println(builder.String(), builder)

	fmt.Println(ValidBraces("([]){}[]"))
	fmt.Println(ExpressionMatter(1, 3, 1))
	fmt.Println(fmt.Sprintf("%05d", 1))
}

func ExpressionMatter(a int, b int, c int) int {
	arr := []int{a, b, c}
	sort.Ints(arr)
	tmp := 0
	if arr[0]*arr[1] < arr[0]+arr[1] {
		tmp = arr[0] + arr[1]
	} else {
		tmp = arr[0] * arr[1]
	}
	fmt.Println(tmp*arr[2], tmp+arr[2])
	if tmp*arr[2] < tmp+arr[2] {
		return tmp + arr[2]
	} else {
		return tmp * arr[2]
	}
}

func Comp(array1 []int, array2 []int) bool {
	for _, v := range array1 {
		sq := v * v
		find := false
		fmt.Println(v, sq)
		for _, vv := range array2 {
			if sq == vv {
				find = true
				break
			}
		}
		if !find {
			return false
		}
	}
	return true
}
