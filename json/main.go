package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *User) Get() {
	u.Name = "hello"
	err := json.Unmarshal([]byte(`{"name":"chen","age":20}`), u)
	if err != nil {
		panic(err)
	}
	fmt.Println(u)
}

func main() {
	u := &User{"world", 30}
	u.Get()
}
