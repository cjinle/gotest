package tutorial

import (
	"fmt"
	"os"
	"time"
)

func Time() {
	t := time.Now()
	fmt.Println(t.Local())
	fmt.Println(t.Unix())
	fmt.Println(os.Getpid())
	fmt.Println(t.String())
}
