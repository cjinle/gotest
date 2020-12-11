package cache

import (
	"log"
	"time"

	"github.com/patrickmn/go-cache"
)

type MyStruct struct {
	Name string
	Sex  int
}

func Example() {
	log.Println("cache example start ... ")

	c := cache.New(1*time.Minute, 0)

	c.Set("foo", "bar", cache.DefaultExpiration)

	v, flg := c.Get("foo")

	log.Printf("v = %v, flg = %v\n", v, flg)
	log.Println(v.(string))

	data := &MyStruct{"Jinle Chen", 1}
	c.Set("data", data, cache.DefaultExpiration)

	x, found := c.Get("data")
	if !found {
		log.Println("data is not found!")
	}

	log.Println(x.(*MyStruct).Name, x.(*MyStruct).Sex)

}
