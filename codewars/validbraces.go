package codewars

// import "fmt"

func ValidBraces(str string) bool {
	s := []byte{}
	m := map[byte]byte{'(':')','[':']','{':'}'}
	for _, v := range []byte(str) {
		if len(s) > 0 && m[s[len(s)-1]] == v {
			s = s[:len(s)-1]
		} else {
			s = append(s, v)
		}
	} 
	return len(s) == 0
}