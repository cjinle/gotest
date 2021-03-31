package mstat2

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/cjinle/game/lib/errs"
	"github.com/cjinle/game/lib/golog"
	"github.com/cjinle/game/lib/util"

	"github.com/gomodule/redigo/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	defaultDatabase   = "mstat"
	defaultCollection = "mstat"
	expire            = 7 // 缓存过期时间
	collectionNum     = 10
)

const (
	// Invalid unknown type
	Invalid int = iota
	// Int type
	Int
	// Float type
	Float
	// String type
	String
)

var intFields = []string{
	"lgdays",   // 登录天数
	"maxwin",   // 最大赢取
	"maxwin_",  // 最大赢取场次
	"maxwon",   // 最大拥有
	"payTimes", // 支付次数

}
var floatFields = []string{
	"payAmount", // 累计支付金额
	"score",
}
var stringFields = []string{
	"msisdn", // 手机号
}

var fields = map[string]reflect.Kind{}

var err error
var logger *golog.Logger

// Mstat 每日用户数据统计结构
type Mstat struct {
	redisPool   *redis.Pool
	mongoClient *mongo.Client
}

// Ints 整型数组
type Ints map[string]int

// Floats 浮点型数组
type Floats map[string]float64

// Strings 字符串数组
type Strings map[string]string

func init() {
	logger = golog.NewLogger(
		&golog.LogConf{
			Path:   os.Getenv("GAMEPATH") + "/log",
			Level:  golog.Debug,
			Prefix: "mstat",
		},
	)
}

// New 初始化对象
func New() *Mstat {
	redisPool := redis.NewPool(func() (redis.Conn, error) {
		con, err := redis.Dial("tcp", "127.0.0.1:6379")
		if err != nil {
			return nil, err
		}
		return con, nil
	}, 10)
	ms := &Mstat{
		redisPool: redisPool,
	}
	ms.mongoClient, err = mongo.Connect(context.Background(),
		options.Client().ApplyURI("mongodb://127.0.0.1:27017/?maxPoolSize=10"))
	if err != nil {
		logger.Fatalf("%v", err)
	}
	return ms
}

// Close 关闭资源
func (ms *Mstat) Close() {
	err = ms.mongoClient.Disconnect(context.Background())
	logger.Infof("close mongo client %v", err)
	err = ms.redisPool.Close()
	logger.Infof("close redis client %v", err)
	logger.Close()
}

// MsetInt 批量设置缓存--整型
func (ms *Mstat) MsetInt(mid uint32, value *map[string]int) error {
	if mid == 0 {
		return errs.MidError
	}
	err = util.FilterInts(value, &intFields)
	if err != nil {
		return err
	}
	con := ms.redisPool.Get()
	defer con.Close()
	cacheKey := ms.getCacheKey(mid, Int)
	_, err = con.Do("HMSET", util.MapInt2RedisArgs(cacheKey, value)...)
	if err != nil {
		con.Do("EXPIRE", cacheKey, expire*86400)
	}
	return err
}

// MsetFloat 批量设置缓存--浮点型
func (ms *Mstat) MsetFloat(mid uint32, value *map[string]float64) error {
	if mid == 0 {
		return errs.MidError
	}
	err = util.FilterFloats(value, &floatFields)
	if err != nil {
		return err
	}
	con := ms.redisPool.Get()
	defer con.Close()
	cacheKey := ms.getCacheKey(mid, Float)
	_, err = con.Do("HMSET", util.MapFloat2RedisArgs(cacheKey, value)...)
	if err != nil {
		con.Do("EXPIRE", cacheKey, expire*86400)
	}
	return err
}

// MsetString 批量设置缓存--浮点型
func (ms *Mstat) MsetString(mid uint32, value *map[string]string) error {
	if mid == 0 {
		return errs.MidError
	}
	err = util.FilterStrings(value, &stringFields)
	if err != nil {
		return err
	}
	con := ms.redisPool.Get()
	defer con.Close()
	cacheKey := ms.getCacheKey(mid, String)
	_, err = con.Do("HMSET", util.MapString2RedisArgs(cacheKey, value)...)
	if err != nil {
		con.Do("EXPIRE", cacheKey, expire*86400)
	}
	return err
}

// MgetInt 批量获取数据
func (ms *Mstat) MgetInt(mid uint32, fieldArr []string) (*map[string]int, error) {
	m := make(map[string]int)
	if mid == 0 {
		return &m, errs.MidError
	}
	con := ms.redisPool.Get()
	defer con.Close()
	cacheKey := ms.getCacheKey(mid, Int)
	if len(fieldArr) == 0 {
		m, err = redis.IntMap(con.Do("HGETALL", util.RedisHMGetCmd(cacheKey, fieldArr)...))
	} else {
		results, err := redis.Ints(con.Do("HMGET", util.RedisHMGetCmd(cacheKey, fieldArr)...))
		if err == nil {
			m = util.RedisHMGetResultInts(fieldArr, results)
		}
	}
	return &m, err
}

// MgetFloat 批量获取数据
func (ms *Mstat) MgetFloat(mid uint32, fieldArr []string) (*map[string]float64, error) {
	m := make(map[string]float64)
	if mid == 0 {
		return &m, errs.MidError
	}
	con := ms.redisPool.Get()
	defer con.Close()
	cacheKey := ms.getCacheKey(mid, Float)
	if len(fieldArr) == 0 {
		results, err := redis.StringMap(con.Do("HGETALL", util.RedisHMGetCmd(cacheKey, fieldArr)...))
		if err != nil {
			for k, v := range results {
				m[k], _ = strconv.ParseFloat(v, 64)
			}
		}
	} else {
		results, err := redis.Float64s(con.Do("HMGET", util.RedisHMGetCmd(cacheKey, fieldArr)...))
		if err == nil {
			m = util.RedisHMGetResultFloats(fieldArr, results)
		}
	}
	return &m, err
}

// IncrBy 整型自增
func (ms *Mstat) IncrBy(mid uint32, field string, num int) (int, error) {
	if mid == 0 {
		return 0, errs.MidError
	}
	if !util.InArrayString(field, &intFields) {
		return 0, errs.FieldNotExist
	}
	con := ms.redisPool.Get()
	defer con.Close()
	cacheKey := ms.getCacheKey(mid, Int)
	num, err = redis.Int(con.Do("HINCRBY", cacheKey, field, num))
	if err != nil {
		return 0, err
	}
	con.Do("EXPIRE", cacheKey, expire*86400)
	return num, err
}

// IncrByFloat 浮点型自增
func (ms *Mstat) IncrByFloat(mid uint32, field string, num float64) (float64, error) {
	if mid == 0 {
		return 0, errs.MidError
	}
	if !util.InArrayString(field, &intFields) {
		return 0, errs.FieldNotExist
	}
	con := ms.redisPool.Get()
	defer con.Close()
	cacheKey := ms.getCacheKey(mid, Float)
	num, err = redis.Float64(con.Do("HINCRBYFLOAT", cacheKey, field, num))
	if err != nil {
		return 0, err
	}
	con.Do("EXPIRE", cacheKey, expire*86400)
	return num, err
}

// SetMax 设置最大值
func (ms *Mstat) SetMax(mid uint32, field string, value int) error {
	if mid == 0 {
		return errs.MidError
	}
	return nil
}

// Cache2DB 缓存落地
func (ms *Mstat) Cache2DB(mid uint32, date string) error {
	return nil
}

// DB2Cache 从DB加载到缓存
func (ms *Mstat) DB2Cache(mid uint32, date string) error {
	return nil
}

// getCollection 获取对应表
func (ms *Mstat) getCollection(mid uint32) *mongo.Collection {
	collection := fmt.Sprintf("%s%x", defaultCollection, mid%collectionNum)
	return ms.mongoClient.Database(defaultDatabase).Collection(collection)
}

func (ms *Mstat) getCacheKey(mid uint32, valueType int) string {
	return fmt.Sprintf("%s:%d:%d", defaultCollection, mid, valueType)
}
