package cache

import (
	"log"
	"time"

	"github.com/patrickmn/go-cache"
)

// MyStruct is test struct
type MyStruct struct {
	Name string
	Sex  int
}

// Example is example
func Example() {
	log.Println("cache example start ... ")

	c := cache.New(1*time.Minute, 10*time.Second)
	c.OnEvicted(func(k string, v interface{}) {
		log.Println("expire callback", k, v)
	})

	c.Set("foo", "bar", cache.DefaultExpiration)

	v, flg := c.Get("foo")

	log.Printf("v = %v, flg = %v\n", v, flg)
	log.Println(v.(string))

	data := &MyStruct{"Jinle Chen", 1}
	c.Set("data", data, 5*time.Second)
	i := 1
	for {
		x, found := c.Get("data")
		if !found {
			log.Println("data is not found!")
			data = &MyStruct{"xxxxxxx", i}
			c.Set("data", data, time.Second)
		} else {
			log.Println(x)
		}
		// log.Println(x.(*MyStruct).Name, x.(*MyStruct).Sex)
		log.Println((*data).Name, data.Sex)
		time.Sleep(1 * time.Second)
		log.Println(i)
		i++
	}

}
