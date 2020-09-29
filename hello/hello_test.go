package hello

import (
	"fmt"
	"testing"
)

func TestPrint(t *testing.T) {
	fmt.Println("test print call")
	fmt.Println(t)
	t.Log(123)
}
