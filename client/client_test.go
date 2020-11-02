package client

import (
	"fmt"
	"testing"
)

func TestSth(t *testing.T) {
	cli, _ := NewClient()
	fmt.Println(cli)
	cli.GetGameInfo(2494)
	fmt.Println("done")
}