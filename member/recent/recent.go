package recent

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/cjinle/game/lib/golog"

	"github.com/patrickmn/go-cache"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	defaultExpire       = 3600
	defaultCleanInetval = 60
	defaultDatabase     = "recent"
	defaultCollection   = "recent"
	collectionNum       = 10
)

// Recent 最近信息
type Recent struct {
	cache        *cache.Cache
	mongoClient  *mongo.Client
	cache2DBChan chan uint32
}

// UserRecent 用户最近信息
type UserRecent struct {
	Mid       uint32 `bson:"mid"`
	API       uint32 `bson:"api"`
	Ver       string `bson:"ver"`
	Lang      uint8  `bson:"lang"`
	Mac       string `bson:"mac"`
	Imei      string `bson:"imei"`
	IP        string `bson:"ip"`
	DeviceID  string `bson:"deviceid"`
	AndroidID string `bson:"androidid"`
}

var err error
var logger *golog.Logger

func init() {
	logger = golog.NewLogger(
		&golog.LogConf{
			Path:   os.Getenv("GAMEPATH") + "/log",
			Level:  golog.Debug,
			Prefix: defaultDatabase,
		},
	)
}

// New 初始化
func New() *Recent {
	recent := &Recent{}
	recent.cache = cache.New(defaultExpire*time.Second, defaultCleanInetval*time.Second)
	recent.mongoClient, err = mongo.Connect(context.Background(),
		options.Client().ApplyURI("mongodb://127.0.0.1:27017/?maxPoolSize=10"))
	if err != nil {
		logger.Fatalf("%v", err)
	}
	recent.cache2DBChan = make(chan uint32, 1000)
	go func(recent *Recent) {
		for {
			select {
			case mid := <-recent.cache2DBChan:
				recent.Cache2DB(mid)
				// logger.Debugf("cache2DBChan pop mid: %d, err: %v", mid, err)
			}
		}
	}(recent)
	return recent
}

// Close 关闭资源
func (recent *Recent) Close() {
	recent.cache.Flush()
	recent.mongoClient.Disconnect(context.Background())
	close(recent.cache2DBChan)
}

// Get 获取用户最近信息
func (recent *Recent) Get(mid uint32) (*UserRecent, error) {
	cacheKey := fmt.Sprint(mid)
	if data, ok := recent.cache.Get(cacheKey); ok {
		return data.(*UserRecent), nil
	}
	ur := &UserRecent{Mid: mid}
	recent.GetCollection(mid).FindOne(context.TODO(), bson.M{"mid": mid}).Decode(ur)
	recent.cache.Set(fmt.Sprint(mid), ur, defaultExpire*time.Second)
	return ur, nil
}

// Cache2DB 缓存落地
func (recent *Recent) Cache2DB(mid uint32) error {
	cacheKey := fmt.Sprint(mid)
	if data, ok := recent.cache.Get(cacheKey); ok {
		ur := data.(*UserRecent)
		opts := options.Update().SetUpsert(true)
		_, err := recent.GetCollection(mid).UpdateOne(context.TODO(),
			bson.M{"mid": mid},
			bson.M{"$set": bson.M{
				"api":       ur.API,
				"ver":       ur.Ver,
				"lang":      ur.Lang,
				"mac":       ur.Mac,
				"imei":      ur.Imei,
				"ip":        ur.IP,
				"deviceid":  ur.DeviceID,
				"androidid": ur.AndroidID,
			}}, opts)
		if err != nil {
			logger.Errorf("cache2db mid: %d", mid)
			return err
		}
	} else {
		return fmt.Errorf("mid: %d not in cache", mid)
	}
	return nil
}

// AddCache2DBChan 缓存落地队列
func (recent *Recent) AddCache2DBChan(mid uint32) {
	// recent.cache2DBChan <- mid
	go func() {
		recent.cache2DBChan <- mid
	}()
}

// GetCollection 获取对应表
func (recent *Recent) GetCollection(mid uint32) *mongo.Collection {
	collection := fmt.Sprintf("%s%d", defaultCollection, mid%collectionNum)
	return recent.mongoClient.Database(defaultDatabase).Collection(collection)
}
