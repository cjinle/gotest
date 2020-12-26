package mredis

import (
	"context"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type MRedis struct {
	con        redis.Conn
	ctx        context.Context
	cancel     context.CancelFunc
	ExpireChan chan RedisKeyExpire
}

type RedisKeyExpire struct {
	Key  string
	Time int32
}

func NewRedis(addr string) *MRedis {
	fmt.Println("NewRedis...", addr)

	con, err := redis.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	r := &MRedis{
		con:        con,
		ExpireChan: make(chan RedisKeyExpire),
	}
	r.ctx, r.cancel = context.WithCancel(context.Background())
	go r.Start()
	return r
}

func (r *MRedis) Start() {
	for {
		select {
		case data := <-r.ExpireChan:
			r.DoExpire(&data)
		case <-r.ctx.Done():
			return

		}
	}

}

func (r *MRedis) Get(key, resultTyte string) (interface{}, error) {
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

func (r *MRedis) Set(key string, value interface{}) (interface{}, error) {
	return r.con.Do("SET", key, value)
}

func (r *MRedis) Ping() (string, error) {
	return redis.String(r.con.Do("PING"))
}

func (r *MRedis) Expire(key string, time int32) {
	r.ExpireChan <- RedisKeyExpire{key, time}
}

func (r *MRedis) DoExpire(data *RedisKeyExpire) {
	fmt.Println("do expire key =", data.Key, ", time =", data.Time)
}

func (r *MRedis) Close() {
	// close(r.ExpireChan)
	r.cancel()
	r.con.Close()
}
