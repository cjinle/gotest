package codewars

import "strings"

func WeirdCase(str string) string {
	bytes := []byte(strings.ToUpper(str))
	i := 0
	for k, v := range bytes {
		if v == byte(' ') {
			i = 0
			continue
		}
		if i % 2 == 1 {
			bytes[k] = v + 32
		}
		i++
	}
	return string(bytes)
}