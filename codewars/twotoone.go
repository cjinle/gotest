package codewars

import (
	// "fmt"
	"sort"
	"strings"
)


func TwoToOne(s1 string, s2 string) string {
	// fmt.Println(s1, s2)
	bytes := []byte(s1 + s2)
	set := make(map[byte]int)
	for _, v := range bytes {
		if _, ok := set[v]; ok {
			set[v] += 1
		} else {
			set[v] = 1
		}
	}
	arr := []int{}
	for k, _ := range set {
		arr = append(arr, int(k))
	}
	sort.Ints(arr)

	out := []byte{}
	for _, v := range arr {
		out = append(out, byte(v))
	}
	return string(out)
}

func TwoToOne2(s1, s2 string) string {
	chars := strings.Split(s1+s2, "")
	sort.Strings(chars)
	result := ""
	for _, s := range chars {
		chr := string(s)
		if !strings.Contains(result, chr) {
			result = result + chr
		}
	}
	return result
}

func TwoToOne3(s1 string, s2 string) string {
  result := ""
  for _, char := range strings.Split("abcdefghijklmnopqrstuvwxyz", "") {
    if strings.Contains(s1+s2, char) {
      result += char
    }
  }
  return result
}
