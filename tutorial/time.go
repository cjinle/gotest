package tutorial

import (
	"fmt"
	"time"
)

func Tick() {
	for range time.Tick(time.Second * 5) {
		fmt.Println("tick...")
	}
}
