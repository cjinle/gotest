package client

import (
	"fmt"
	"testing"
)

func TestSth(t *testing.T) {
	cli, _ := NewClient()
	defer cli.Close()
	fmt.Println(cli)

	fmt.Println(cli.UpdateMoney(2494, 1, 100))

	ret := cli.GetGameInfo(2494)
	info, ok := ret.(map[string]interface{})
	if ok {
		fmt.Println(info["exp"])
		fmt.Println(info["level"])
		fmt.Println(info["money"])
		fmt.Println(info["coin"])
		fmt.Println(info["attr1"])
		fmt.Println(info["safebox"])
	}
	
	fmt.Println("done")
}