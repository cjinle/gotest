package mylib

import (
	"fmt"
)

type I interface {
	Info() string
}

type Person struct {
	Name string
	Sex  int
}

func (p *Person) Info() string {
	p.Name = "xxx"
	fmt.Println(p.Name)
	return fmt.Sprintf("name: %v, sex: %d", p.Name, p.Sex)
}

func InterfaceOuput() {
	var i I = &Person{"chenjinle", 1}
	fmt.Println(i.Info())
	fmt.Println(i)
}
