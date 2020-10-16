package tutorial

import (
	"fmt"
	"os"
	"time"
)

// 文件信息
func Stat() {
	var file string
	file = "aa"

	fi, err := os.Stat(file)

	fmt.Println(fi.Name(), fi.IsDir(), fi.Mode().String(), err)
}


// ----- time -------
func Time() {
	t := time.Now()
	fmt.Println(t.Local())
	fmt.Println(t.Unix())
	fmt.Println(os.Getpid())
	fmt.Println(t.String())
}
