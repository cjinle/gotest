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
	Redis         string      `yaml:"redis"`
	CacheKey      string      `yaml:"cache-key"`
	AndroidWorker int         `yaml:"android-worker"`
	IOSWorker     int         `yaml:"ios-worker"`
	API           map[int]API `yaml:"api"`
}

// Push data struct
type Push struct {
	// redis     redis.Conn
	redisPool *redis.Pool
	config    *Config
}

// Notification 推送内容
type Notification struct {
	API     int      `json:"api"`
	Tokens  []string `json:"tokens"`
	Content string   `json:"content"`
}

// AndroidClient 安卓推送客户端
type AndroidClient struct {
	ID int
}

// IOSClient IOS推送客户端
type IOSClient struct {
	ID int
}

const (
	maxWorkNum = 50
)

var err error
var androidClients chan *AndroidClient
var iosClients chan *IOSClient

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
	push.redisPool = redis.NewPool(func() (redis.Conn, error) {
		con, err := redis.Dial("tcp", push.config.Redis)
		if err != nil {
			return nil, err
		}
		return con, nil
	}, 2)

	go push.StartWorkers()
	go push.Pop()
	go push.AddTestData()
	return push
}

// Start 开启推送
func (push *Push) Start() {
	select {}
}

// StartWorkers 开启处理队列worker
func (push *Push) StartWorkers() {
	androidClients = make(chan *AndroidClient, maxWorkNum)
	if push.config.AndroidWorker > maxWorkNum {
		push.config.AndroidWorker = maxWorkNum
	}
	for i := 0; i < push.config.AndroidWorker; i++ {
		androidClients <- &AndroidClient{ID: i}
	}

	iosClients = make(chan *IOSClient, maxWorkNum)
	if push.config.IOSWorker > maxWorkNum {
		push.config.IOSWorker = maxWorkNum
	}
	for i := 0; i < push.config.IOSWorker; i++ {
		iosClients <- &IOSClient{ID: i}
	}
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

func genJSON() ([]byte, error) {
	notice := &Notification{
		API:     0x01001200,
		Tokens:  []string{"aaaaa", "bbbb", "ccc"},
		Content: "test message...",
	}
	return json.Marshal(notice)
}
