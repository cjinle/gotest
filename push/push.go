package push

import (
	"encoding/json"
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
	// redis     redis.Conn
	redisPool *redis.Pool
	config    *Config
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
	// conn, err := redis.Dial("tcp", push.config.Redis)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// push.redis = conn
	push.redisPool = redis.NewPool(func() (redis.Conn, error) {
		con, err := redis.Dial("tcp", "127.0.0.1:6379")
		if err != nil {
			return nil, err
		}
		return con, nil
	}, 2)

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
	conn := push.redisPool.Get()
	defer conn.Close()
	for {
		str, err := redis.Bytes(conn.Do("RPOP", push.config.CacheKey))

		if err != nil {
			pong, _ := conn.Do("PING")
			log.Println(pong, err)
			time.Sleep(1 * time.Second)
			continue
		}
		log.Println("--->>>", string(str), time.Now().Unix())
		time.Sleep(10000 * time.Microsecond)
	}
}

// AddTestData 加数据
func (push *Push) AddTestData() {
	conn := push.redisPool.Get()
	defer conn.Close()
	rand.Seed(time.Now().UnixNano())
	i := 0
	for {
		value, _ := genJSON()
		log.Println(value)
		conn.Do("LPUSH", push.config.CacheKey, string(value))
		i++
		if i > 10000 {
			time.Sleep(time.Minute)
			i = 0
			continue
		}
		log.Println("<<<---", value)
		time.Sleep(10000 * time.Microsecond)
	}
}

type pushData struct {
	TermID    string `json:"termid"`
	ContentID string `json:"contentid"`
	Token     string `json:"token"`
}

func genJSON() ([]byte, error) {
	return json.Marshal(&pushData{TermID: "2020-12-31", ContentID: "yesterday", Token: "xxxxxx"})
}
