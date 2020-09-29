package tutorial

import (
	"fmt"
)

type I2 interface {
	Info() string
	AgeInfo() string
}

type Person2 struct {
	Name string
	Age  int
}

func (p *Person2) String() string {
	return fmt.Sprintf("[string]my name: %s, age %d", p.Name, p.Age)
}

func (p *Person2) Info() string {
	return fmt.Sprintf("my name: %s, age %d", p.Name, p.Age)
}

func (p *Person2) AgeInfo() string {
	return fmt.Sprintf("age %d", p.Age)
}

func Interface2Output() {
	var i I2
	person := &Person2{"chenjinle", 18}
	i = person
	fmt.Println(person)
	fmt.Println(i.Info())
	fmt.Println(i.AgeInfo())

	i = &Person2{"zhangsan", 100}
	fmt.Println(i)
	fmt.Println(i.Info())
	fmt.Println(i.AgeInfo())

}
