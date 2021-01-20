package uuid

import (
	"fmt"
	"testing"
)

func TestOne(t *testing.T) {
	uid := NewUUID()
	fmt.Println(uid)

	NewGoogleUUID()

	fmt.Println(NewObjectID())
}
