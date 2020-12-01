package tutorial

import (
	"encoding/json"
	"fmt"
	"log"
)

// ----------- array ------------
// array example
func ArrOuput() {
	var a = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a)

	r := [...]int{99: -1}
	fmt.Println(r)

	fmt.Println("hello world")

	// var str = [...]byte{'a', 'b', 'c'}
	// str = append(str[:], 'x')
	// fmt.Println(str)
}

// --------- slice ---------
func Slice() {
	s := []string{"a", "b", "c"}
	fmt.Println(s)
	s = s[0:0]
	fmt.Println(s)
}

// --------------- map --------
// map example
type La struct {
	x, y int
}

func Map() {
	var m = map[string]La{
		"a": {1, 2},
		"b": {3, 4},
	}

	fmt.Println(m["a"].x + m["b"].y)
}

// ------------- make ----------------
// map create example
type UserMap struct {
	uid   int
	money int
}

func Make() {
	usermap := make([]int, 10)
	fmt.Println(usermap)
}

// ------------- struct ------------
type S struct {
	int
	string
}

func Struct2() {
	s := S{1, "hello"}
	fmt.Println(s)
	fmt.Println(s.int)
	fmt.Println(s.string)
}

// ------------- interface ----------------
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

// --------------------------

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

// ------------ json ---------------
type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func MyJson() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}

	data, err := json.Marshal(movies)
	log.Printf("JSON Data: %s", data)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	str := `{Title: "hello", year: 2020, Color: true, Actor: ["John", "Jim"]}`
	// var v Movie
	// v := &Movie{}
	var v interface{}
	v = v.(Movie)
	err = json.Unmarshal([]byte(str), v)
	fmt.Println("\n", str, v)

}
