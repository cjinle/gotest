package tutorial

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func Redis() {
	fmt.Println("redis connect")
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.100.107:6379",
		Password: "",
		DB:       0,
	})
	ret, err := client.Ping().Result()
	fmt.Println(ret, err)

	cacheKey := "go_test"

	// set cache and time expire
	err = client.Set(cacheKey, "111111", 86400*time.Second).Err()
	if err != nil {
		panic(err)
	}
	// get cache
	val, err := client.Get(cacheKey).Result()
	if err == redis.Nil {
		fmt.Println("cache key not exist!")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println(val)
	}

	cacheKey = "go_hash"
	fmt.Println(client.HSet(cacheKey, "a", 1).Err())
	fmt.Println(client.HSet(cacheKey, "b", 2).Err())
	fmt.Println(client.HSet(cacheKey, "c", 3).Err())
	fmt.Println(client.HSet(cacheKey, "d", 4).Err())
	valHash, _ := client.HGetAll(cacheKey).Result()
	fmt.Println(valHash)
	fmt.Println(client.HMSet(cacheKey, map[string]interface{}{"c": 4, "d": 5, "e": 6}).Err())
	valHash2, _ := client.HGet(cacheKey, "c").Result()
	fmt.Println(valHash2)
	fmt.Println(client.HIncrBy(cacheKey, "c", 1).Val())

}
