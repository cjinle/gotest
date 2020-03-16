package main


import (
	"fmt"
)

type I interface {
	Info() string
	AgeInfo() string
}

type Person struct {
	Name string
	Age int
}

func (p *Person)String() string {
	return fmt.Sprintf("[string]my name: %s, age %d", p.Name, p.Age)
}

func (p *Person)Info() string {
	return fmt.Sprintf("my name: %s, age %d", p.Name, p.Age)
}

func (p *Person)AgeInfo() string {
	return fmt.Sprintf("age %d", p.Age)
}

func main() {
	var i I
	person := &Person{"chenjinle", 18}
	i = person
	fmt.Println(person)
	fmt.Println(i.Info())
	fmt.Println(i.AgeInfo())

	i = &Person{"zhangsan", 100}
	fmt.Println(i)
	fmt.Println(i.Info())
	fmt.Println(i.AgeInfo())

}

