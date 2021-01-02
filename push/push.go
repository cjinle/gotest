package push

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
	"gopkg.in/yaml.v2"
)

const (
	maxWorkNum  = 50
	androidType = 2
	iosType     = 3
)

// API 推送KEY配置
type API struct {
	Key    string `yaml:"key"`
	Worker int    `yaml:"worker"`
}

// Config 配置
type Config struct {
	Redis       string      `yaml:"redis"`
	CacheKeyPre string      `yaml:"cache-key-pre"`
	TimeOut     int         `yaml:"timeout"`
	API         map[int]API `yaml:"api"`
}

// Push data struct
type Push struct {
	// redis     redis.Conn
	redisPool *redis.Pool
	config    *Config
}

// Client 接口
type Client interface {
	Send(*Message) error
	Close()
}

// AndroidClient 安卓推送客户端
type AndroidClient struct {
	ID  int
	API int
	Key string
}

// IOSClient IOS推送客户端
type IOSClient struct {
	ID  int
	API int
	Key string
}

// Message 推送内容
type Message struct {
	API     int      `json:"api"`
	Tokens  []string `json:"tokens"`
	Content string   `json:"content"`
}

var err error
var messageChan map[int]chan *Message
var clientChan map[int]chan interface{}

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
	}, 10)

	go push.AddTestData()
	return push
}

// Start 开启推送
func (push *Push) Start() {
	log.Println("push start ... ")
	messageChan = make(map[int]chan *Message)
	clientChan = make(map[int]chan interface{})

	// 初始化client队列
	for api, cfg := range push.config.API {
		log.Println("range cfg ", api, cfg)
		messageChan[api] = make(chan *Message, 1000)
		clientChan[api] = make(chan interface{}, 50)
		apiType := getClientType(api)
		i := 0
		for i < cfg.Worker {
			if apiType == androidType {
				clientChan[api] <- NewAndroidClient(i, api, cfg.Key)
			} else if apiType == iosType {
				clientChan[api] <- NewIOSClient(i, api, cfg.Key)
			}
			i++
		}
		go push.HandleSend(api, apiType)
		go push.Pop(api)
	}
	select {}
}

// HandleSend 开启处理队列worker
func (push *Push) HandleSend(api int, apiType int) {
	log.Println("HandleSend start ... ")
	var wg sync.WaitGroup
	for {
		client := <-clientChan[api]
		clientChan[api] <- client
		wg.Add(1)
		go func() {
			msg := <-messageChan[api]
			if apiType == androidType {
				client.(*AndroidClient).Send(msg)
			} else {
				client.(*IOSClient).Send(msg)
			}
			wg.Done()
		}()
		wg.Wait()
	}
}

// Pop 从缓存读取任务数据
func (push *Push) Pop(api int) {
	log.Printf("api %x pop start ... ", api)
	if _, ok := messageChan[api]; !ok {
		log.Printf("chan %x not exist!", api)
		return
	}
	conn := push.redisPool.Get()
	defer conn.Close()
	for {
		cacheKey := fmt.Sprintf("%s:%d", push.config.CacheKeyPre, api)
		str, err := redis.Bytes(conn.Do("RPOP", cacheKey))
		if err != nil {
			log.Printf("redis message %x %v", api, err)
			conn.Do("PING")
			time.Sleep(time.Duration(rand.Intn(30)) * time.Second)
			continue
		}
		var msg Message
		err = json.Unmarshal(str, &msg)
		if err != nil {
			log.Printf("redis message json decode %x %v", api, err)
		}
		messageChan[api] <- &msg
		// log.Println("message <<< --- ", string(str))
	}
}

// AddTestData 加数据
func (push *Push) AddTestData() {
	conn := push.redisPool.Get()
	defer conn.Close()
	var apis []int
	for api := range push.config.API {
		apis = append(apis, api)
	}
	rand.Seed(time.Now().UnixNano())
	i := 0
	for {
		randNum := rand.Intn(1000000)
		api := apis[(randNum % len(apis))]
		value, _ := genJSON(api)
		cacheKey := fmt.Sprintf("%s:%d", push.config.CacheKeyPre, api)
		conn.Do("LPUSH", cacheKey, string(value))
		i++
		if i > 10000 {
			time.Sleep(time.Minute)
			i = 0
			continue
		}
		// log.Println("<<<---", string(value))
		time.Sleep(1000 * time.Microsecond)
	}
}

func genJSON(api int) ([]byte, error) {
	msg := &Message{
		API:     api,
		Tokens:  []string{"aaaaa", "bbbb", "ccc"},
		Content: "test message...",
	}
	return json.Marshal(msg)
}

func getClientType(api int) int {
	return (api & 0x00000F00) >> 8
}

// NewAndroidClient 创建一个安卓推送对象
func NewAndroidClient(id int, api int, key string) *AndroidClient {
	return &AndroidClient{ID: id, API: api, Key: key}
}

// Send 安卓推送
func (client *AndroidClient) Send(msg *Message) error {
	log.Printf("client: %x %d start send ... ", client.API, client.ID)
	log.Println("send msg >>> ", msg)
	time.Sleep(time.Duration(rand.Intn(10000)) * time.Microsecond)
	log.Printf("client: %d send done.", client.ID)
	return nil
}

// Close 关闭连接
func (client *AndroidClient) Close() {
	log.Printf("client: %d close", client.ID)
}

// NewIOSClient 创建一个安卓推送对象
func NewIOSClient(id int, api int, key string) *IOSClient {
	return &IOSClient{ID: id, API: api, Key: key}
}

// Send IOS推送
func (client *IOSClient) Send(msg *Message) error {
	log.Printf("client: %x %d start send ... ", client.API, client.ID)
	log.Println("send msg >>> ", msg)
	time.Sleep(time.Duration(rand.Intn(10000)) * time.Microsecond)
	log.Printf("client: %d send done.", client.ID)
	return nil
}

// Close 关闭连接
func (client *IOSClient) Close() {
	log.Printf("client: %d close", client.ID)
}
