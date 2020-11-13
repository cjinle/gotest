package main

import (
	"fmt"
	"os"
)

type Fd interface{}

func NewFile() (Fd, err error) {
	return os.OpenFile("xx.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
}

func main() {
	fmt.Println("test main starting ... ")
	fd, _ = NewFile()
	fd.logs("a")
	fd.logs("b")
}

func (fd *Fd) logs(str string) {
	buf := []byte(str + "\n")
	fd.Write(buf)
}
