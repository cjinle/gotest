package utils

import (
	"fmt"
	"testing"
)

func TestOne(t *testing.T) {
	fmt.Println(Md5Sum("123456"))
}