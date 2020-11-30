package uuid

import (
	"fmt"
	"testing"
)

func TestOne(t *testing.T) {
	uid := New()
	fmt.Println(uid)
}