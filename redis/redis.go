package redis

import (
	"context"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type Redis struct {
	con        redis.Conn
	ctx        context.Context
	cancel     context.CancelFunc
	ExpireChan chan RedisKeyExpire
}

type RedisKeyExpire struct {
	Key  string
	Time int32
}

func NewRedis(addr string) *Redis {
	fmt.Println("NewRedis...", addr)

	con, err := redis.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	r := &Redis{
		con:        con,
		ExpireChan: make(chan RedisKeyExpire),
	}
	r.ctx, r.cancel = context.WithCancel(context.Background())
	go r.Start()
	return r
}

func (r *Redis) Start() {
	for {
		select {
		case data := <-r.ExpireChan:
			r.DoExpire(&data)
		case <-r.ctx.Done():
			return

		}
	}

}

func (r *Redis) Get(key, resultTyte string) (interface{}, error) {
	resp, err := r.con.Do("GET", key)
	if err != nil {
		return resp, err
	}
	switch resultTyte {
	case "int":
		return redis.Int(resp, err)
	case "int64":
		return redis.Int64(resp, err)
	case "uint64":
		return redis.Uint64(resp, err)
	case "float64":
		return redis.Float64(resp, err)
	case "string":
		return redis.String(resp, err)
	default:
		return redis.String(resp, err)
	}
}

func (r *Redis) Set(key string, value interface{}) (interface{}, error) {
	return r.con.Do("SET", key, value)
}

func (r *Redis) Ping() (string, error) {
	return redis.String(r.con.Do("PING"))
}

func (r *Redis) Expire(key string, time int32) {
	r.ExpireChan <- RedisKeyExpire{key, time}
}

func (r *Redis) DoExpire(data *RedisKeyExpire) {
	fmt.Println("do expire key =", data.Key, ", time =", data.Time)
}

func (r *Redis) Close() {
	// close(r.ExpireChan)
	r.cancel()
	r.con.Close()
}

type RedisPool *redis.Pool

func NewRedisPool() *redis.Pool {
	redisPool := redis.NewPool(func() (redis.Conn, error) {
		con, err := redis.Dial("tcp", "127.0.0.1:6379")
		if err != nil {
			return nil, err
		}
		return con, nil
	}, 10)
	// redisPool.MaxActive = 200
	// redisPool.Wait = true
	// redisPool.IdleTimeout = 240 * time.Second
	return redisPool
}
