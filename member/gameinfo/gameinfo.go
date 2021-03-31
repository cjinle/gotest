package gameinfo

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/cjinle/game/lib/errs"
	"github.com/cjinle/game/lib/golog"
	"github.com/patrickmn/go-cache"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	defaultExpire       = 3600
	defaultCleanInetval = 60
	defaultDatabase     = "member"
	defaultCollection   = "gameinfo"
)

// GameInfo 游戏数据结构
type GameInfo struct {
	cache           *cache.Cache
	mongoCollection *mongo.Collection
	mongoClient     *mongo.Client
}

var cache2DBChan chan uint32
var logger *golog.Logger
var err error

func init() {
	logger = golog.NewLogger(
		&golog.LogConf{
			Path:   os.Getenv("GAMEPATH") + "/log",
			Level:  golog.Debug,
			Prefix: "gameinfo",
		},
	)
}

// New 初始化一个GameInfo对象
func New() *GameInfo {
	gi := &GameInfo{}
	gi.cache = cache.New(defaultExpire*time.Second, defaultCleanInetval*time.Second)
	gi.mongoClient, err = mongo.Connect(context.Background(),
		options.Client().ApplyURI("mongodb://127.0.0.1:27017/?maxPoolSize=10"))
	if err != nil {
		logger.Fatalf("%v", err)
	}
	gi.mongoCollection = gi.mongoClient.Database(defaultDatabase).Collection(defaultCollection)
	cache2DBChan = make(chan uint32, 1000)
	go func(gi *GameInfo) {
		for {
			select {
			case mid := <-cache2DBChan:
				err := gi.Cache2DB(mid)
				if err != nil {

				}
				// log.Println(">>>> cache2DBChan", mid, err)
			}
		}

	}(gi)
	return gi
}

// GetCache 获取用户缓存数据
func (gi *GameInfo) GetCache(mid uint32) (*UserGameInfo, error) {
	cacheKey := fmt.Sprint(mid)
	if data, ok := gi.cache.Get(cacheKey); ok {
		return data.(*UserGameInfo), nil
	}

	// load db to cache
	ugi := &UserGameInfo{}
	err := gi.mongoCollection.FindOne(context.TODO(), bson.M{"mid": mid}).Decode(ugi)
	if err != nil {
		return &UserGameInfo{}, errs.RecordNotExist
	}

	err = gi.SetCache(mid, ugi)
	if err != nil {
		log.Println(err)
	}
	return ugi, nil
}

// SetCache 更新用户信息到缓存
func (gi *GameInfo) SetCache(mid uint32, ugi *UserGameInfo) error {
	if mid != ugi.Mid {
		return fmt.Errorf("mid: %d set cache error", mid)
	}
	gi.cache.Set(fmt.Sprint(mid), ugi, defaultExpire*time.Second)
	return nil
}

// Cache2DB 缓存数据回到DB
func (gi *GameInfo) Cache2DB(mid uint32) error {
	cacheKey := fmt.Sprint(mid)
	if data, ok := gi.cache.Get(cacheKey); ok {
		ugi := data.(*UserGameInfo)
		// log.Println("==== Cache2DB ====", ugi)
		_, err := gi.mongoCollection.UpdateOne(context.TODO(),
			bson.M{"mid": mid},
			bson.M{"$set": bson.M{
				"money":   ugi.Money,
				"coin":    ugi.Coin,
				"exp":     ugi.Exp,
				"level":   ugi.Level,
				"safebox": ugi.Safebox,
			}})
		ugi.InChan = false
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("Mid: %d not in cache", mid)
	}
	return nil
}

// CreateUserGameInfo 创建用户信息
func (gi *GameInfo) CreateUserGameInfo(mid uint32) (*UserGameInfo, error) {
	ugi := newUserGameInfo(mid)
	_, err := gi.mongoCollection.InsertOne(context.TODO(), bson.M{"mid": mid})
	return ugi, err
}

// Close 关闭资源
func (gi *GameInfo) Close() {
	gi.mongoClient.Disconnect(context.Background())
	logger.Close()
}

// UserGameInfo 用户游戏数据结构
type UserGameInfo struct {
	Mid     uint32 `bson:"mid"`
	Money   uint32 `bson:"money"`
	Coin    uint32 `bson:"coin"`
	Exp     uint32 `bson:"exp"`
	Level   uint8  `bson:"level"`
	Safebox uint32 `bson:"safebox"`
	Mu      sync.Mutex
	InChan  bool
}

func newUserGameInfo(mid uint32) *UserGameInfo {
	return &UserGameInfo{Mid: mid}
}

// String 用户信息输出
func (ugi *UserGameInfo) String() string {
	return fmt.Sprintf("Mid: %d, Money: %d, Coin: %d, Exp: %d, Level: %d, Safebox: %d",
		ugi.Mid, ugi.Money, ugi.Coin, ugi.Exp, ugi.Level, ugi.Safebox)
}

// AddMoney 加减筹码
func (ugi *UserGameInfo) AddMoney(mode uint32, num int32) error {
	if num == 0 {
		return nil
	}
	ugi.Mu.Lock()
	defer ugi.Mu.Unlock()
	if num < 0 {
		if ugi.Money < uint32(num*-1) {
			return fmt.Errorf("mid: %d, mode: %d, num: %d didn't have enough money", ugi.Mid, mode, num)
		}
		ugi.Money -= uint32(num * -1)
	} else {
		ugi.Money += uint32(num)
	}
	ugi.AddChan()
	return nil
}

// AddCoin 加减金币
func (ugi *UserGameInfo) AddCoin(mode uint32, num int32) error {
	if num == 0 {
		return nil
	}
	ugi.Mu.Lock()
	defer ugi.Mu.Unlock()
	if num < 0 {
		if ugi.Coin < uint32(num*-1) {
			return fmt.Errorf("mid: %d, mode: %d, num: %d didn't have enough coin", ugi.Mid, mode, num)
		}
		ugi.Coin -= uint32(num * -1)
	} else {
		ugi.Coin += uint32(num)
	}
	ugi.AddChan()
	return nil
}

// AddExp 加减经验，自动更新等级
func (ugi *UserGameInfo) AddExp(num int32) error {
	if num == 0 {
		return nil
	}
	ugi.Mu.Lock()
	defer ugi.Mu.Unlock()
	if num < 0 {
		if ugi.Exp < uint32(num*-1) {
			return fmt.Errorf("mid: %d, num: %d didn't have enough exp", ugi.Mid, num)
		}
		ugi.Exp -= uint32(num * -1)
	} else {
		ugi.Exp += uint32(num)
	}
	level := getLevel(ugi.Exp)
	if level != ugi.Level {
		ugi.Level = level
	}
	ugi.AddChan()
	return nil
}

// SetLevel 设置等级
func (ugi *UserGameInfo) SetLevel(mid uint32, level uint8) error {
	if level <= 0 {
		return fmt.Errorf("mid: %d set level error", mid)
	}
	ugi.Level = level
	return nil
}

// AddChan 加到缓存队列
func (ugi *UserGameInfo) AddChan() {
	if !ugi.InChan {
		ugi.InChan = true
		cache2DBChan <- ugi.Mid
	}
}

func getLevel(exp uint32) uint8 {
	levelCfg := []uint32{
		0, 10, 60, 120, 200, 300, 500, 800, 1200, 1700,
		2300, 3100, 4100, 5300, 7800, 10800, 14300, 18300, 23000, 28500,
		35000, 42500, 51000, 60500, 71000, 83000, 97000, 113000, 131000, 151000,
		173000, 197000, 223000, 251000, 281000, 313000, 347000, 383000, 421000, 461000,
		503500, 548500, 596000, 646000, 698500, 753500, 811000, 871000, 934000, 1000000,
	}
	var level uint8
	for idx, needExp := range levelCfg {
		level = uint8(idx + 1)
		if exp < needExp {
			return level
		}
	}
	return level
}
