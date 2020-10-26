package codewars

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func CrackPin(hash string) string {
	for i := 0; i < 100000; i++ {
		pwd := fmt.Sprintf("%05d", i)
		bytes := md5.Sum([]byte(pwd))

		if hash == hex.EncodeToString(bytes[:]) {
			return pwd
		}
	}
	return ""

}

func CrackPin2(hash string) string {
	hashBytesSlice, _ := hex.DecodeString(hash)
	var hashBytes [16]byte
	copy(hashBytes[:], hashBytesSlice)
	for i := 0; i <= 99999; i++ {
		s := fmt.Sprintf("%05d", i)
		if md5.Sum([]byte(s)) == hashBytes {
			return s
		}
	}
	panic("password not found")
}
