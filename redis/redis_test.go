package redis

import (
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = NewRedisPool()
	log.SetFlags(log.Lshortfile)
}

func TestConn(t *testing.T) {
	fmt.Println(".....")

	// a := 23123
	a := "123.12"
	var b interface{} = a
	v := reflect.ValueOf(b)
	fmt.Println(v.Type())

	r := NewRedis("127.0.0.1:6379")
	fmt.Println(r.Get("foo", "string"))
	fmt.Println(r.Get("foo2", "string"))
	fmt.Println(r.Get("foo3", "int64"))
	fmt.Println(r.Set("foo3", "123123123"))
	result, err := r.Get("foo4", "int64")
	if err != nil {
		// panic(err)
		fmt.Println(err)
	}
	// result.(int) += 100
	fmt.Printf("result type = %T, value = %v \n", result, result)
	fmt.Println(r.Set("foo3", 888))
	// r.Close()
	// for {
	// 	select {
	// 	case <-time.After(time.Second * 1):
	// 		fmt.Println(r.Ping())
	// 		// return
	// 	}
	// }

}

func TestPool(t *testing.T) {
	// pool := NewRedisPool()
	con := pool.Get()
	defer con.Close()
	for i := 0; i < 100000; i++ {
		reply, err := con.Do("GET", "foo")
		if err != nil {
			log.Fatal(err)
		}
		log.Println("reply", i, reply, pool.Stats())
	}
	// pool.Close()
	// for {
	log.Println(pool.Stats())
	// 	time.Sleep(time.Second)
	// }
}

func BenchmarkPool(b *testing.B) {
	log.Println(pool.Stats())
	con := pool.Get()
	defer con.Close()
	for i := 0; i < b.N; i++ {
		// reply, err := pool.Get().Do("SET", "foo", rand.Intn(9999999))
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// log.Println("reply", reply)
		reply, err := con.Do("GET", "foo")
		if err != nil {
			log.Fatal(err)
		}
		_, err = redis.Int(reply, err)
		// log.Println("get reply", err, pool.Stats())
	}
	log.Println(pool.Stats())

}
