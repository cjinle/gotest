package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.Local())
	fmt.Println(t.Unix())
	fmt.Println(os.Getpid())
	fmt.Println(t.String())
}
