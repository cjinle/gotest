package sysla

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

type AppConf struct {
	RedisIp   string        `json:redisip`
	RedisPort int           `json:redisport`
	Host      string        `json:host`
	Duration  time.Duration `json:duration`
}

func Run() {
	fmt.Println("sysla start ... ")
	var cfg AppConf
	b, err := ioutil.ReadFile("conf/app.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &cfg)
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.RedisIp, cfg.RedisPort),
		Password: "",
		DB:       0,
	})
	for {
		load, err := GetLoadAvg()
		if err != nil {
			fmt.Println(err)
		}
		now := time.Now()
		year, month, day := now.Date()
		cacheKey := fmt.Sprintf("SYSLA_%s_%d%02d%02d", cfg.Host, year, month, day)
		hour, minute, _ := now.Clock()
		field := fmt.Sprintf("%02d%02d", hour, minute)
		_, err = client.HSet(cacheKey, field, load).Result()
		fmt.Println(field, load)
		time.Sleep(cfg.Duration * time.Second)
	}
}

func GetLoadAvg() (string, error) {
	line, err := ioutil.ReadFile("/proc/loadavg")
	if err != nil {
		return "", err
	}
	values := strings.Fields(string(line))

	load1, err := strconv.ParseFloat(values[0], 64)
	if err != nil {
		return "", err
	}
	load5, err := strconv.ParseFloat(values[1], 64)
	if err != nil {
		return "", err
	}
	load15, err := strconv.ParseFloat(values[2], 64)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v,%v,%v", load1, load5, load15), nil
}
