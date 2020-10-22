package codewars

import (
	"strings"
)

func WeirdCase(str string) string {
	bytes := []byte(strings.ToUpper(str))
	i := 0
	for k, v := range bytes {
		if v == byte(' ') {
			i = 0
			continue
		}
		if i%2 == 1 {
			bytes[k] = v + 32
		}
		i++
	}
	return string(bytes)
}

func WeirdCase2(str string) string {
	bytes := []byte(str)
	i, n := 0, 0
	for k, v := range bytes {
		if v == byte(' ') {
			i = 0
			continue
		}
		switch {
		case i%2 == 0 && v >= byte('a') && v <= byte('z'):
			n = -1
		case i%2 == 1 && v >= byte('A') && v <= byte('Z'):
			n = 1
		default:
			n = 0
		}
		bytes[k] = byte(int(v) + n*32)
		i++
	}
	return string(bytes)
}
