package push

import (
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"github.com/gomodule/redigo/redis"
	"gopkg.in/yaml.v2"
)

// API 推送KEY配置
type API struct {
	Key     string `yaml:"key"`
	TaskNum int    `yaml:"tasknum"`
}

// Config 配置
type Config struct {
	Redis    string      `yaml:"redis"`
	CacheKey string      `yaml:"cache-key"`
	API      map[int]API `yaml:"api"`
}

// Push data struct
type Push struct {
	redis  redis.Conn
	config *Config
}

var err error

func init() {
	// file, err := os.OpenFile("push.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// log.SetOutput(file)
}

// New 初始化
func New() *Push {
	push := &Push{}
	str, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(str, &push.config)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := redis.Dial("tcp", push.config.Redis)
	if err != nil {
		log.Fatal(err)
	}
	push.redis = conn
	go push.Pop()
	go push.AddTestData()
	return push
}

// Start 开启推送
func (push *Push) Start() {
	select {}
}

// Pop 从缓存读取任务数据
func (push *Push) Pop() {
	for {
		str, err := redis.String(push.redis.Do("RPOP", push.config.CacheKey))

		if err != nil {
			pong, _ := push.redis.Do("PING")
			log.Println(pong, err)
			time.Sleep(1 * time.Second)
			continue
		}
		log.Println(str, time.Now().Unix())
		time.Sleep(500 * time.Microsecond)
	}
}

// AddTestData 加数据
func (push *Push) AddTestData() {
	rand.Seed(time.Now().UnixNano())
	i := 0
	for {
		push.redis.Do("LPUSH", push.config.CacheKey, rand.Intn(9999999))
		i++
		if i > 10000 {
			time.Sleep(time.Minute)
			i = 0
			continue
		}
		time.Sleep(500 * time.Microsecond)
	}
}
