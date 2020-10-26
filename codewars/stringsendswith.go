package codewars

import "strings"

func StringsEndsWith(str, ending string) bool {
	if len(str) < len(ending) {
		return false
	}
	return str[len(str)-len(ending):] == ending
}

func StringsEndsWith2(str, ending string) bool {
	return strings.HasSuffix(str, ending)
}
