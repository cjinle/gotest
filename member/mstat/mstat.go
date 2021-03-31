package mstat

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/cjinle/game/lib/errs"
	"github.com/cjinle/game/lib/golog"
	"github.com/cjinle/game/lib/util"
	"github.com/patrickmn/go-cache"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	defaultExpire        = 3600
	defaultCleanInetval  = 60
	defaultDatabase      = "mstat"
	defaultCollection    = "mstat"
	defaultCollectionNum = 10
	expire               = 7 // 缓存过期时间
)

// Mstat 每日用户数据统计结构
type Mstat struct {
	cache       *cache.Cache
	mongoClient *mongo.Client
}

var cache2DBChan chan uint32
var err error
var logger *golog.Logger

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
	ms := &Mstat{}
	ms.cache = cache.New(defaultExpire*time.Second, defaultCleanInetval*time.Second)
	ms.cache.OnEvicted(func(key string, v interface{}) {
		err = ms.Cache2DB(v.(*UserMstat))
		if err != nil {
			logger.Errorf("cache.OnEvicted error mid: %s, %v %v", key, v, err)
		}
	})
	ms.mongoClient, err = mongo.Connect(context.Background(),
		options.Client().ApplyURI("mongodb://127.0.0.1:27017/?maxPoolSize=10"))
	if err != nil {
		logger.Fatalf("%v", err)
	}
	cache2DBChan = make(chan uint32, 1000)
	go func(ms *Mstat) {
		for {
			select {
			case mid := <-cache2DBChan:
				cacheKey := fmt.Sprint(mid)
				if data, ok := ms.cache.Get(cacheKey); ok {
					err = ms.Cache2DB(data.(*UserMstat))
					if err != nil {
						logger.Errorf("cache2DBChan error mid: %d, %v %v", mid, data, err)
					}
				}
			}
		}
	}(ms)
	return ms
}

// GetCache 获取用户数据对象
func (ms *Mstat) GetCache(mid uint32) (*UserMstat, error) {
	cacheKey := fmt.Sprint(mid)
	if data, ok := ms.cache.Get(cacheKey); ok {
		return data.(*UserMstat), nil
	}
	umsd := &UserMstat{Mid: mid}
	err := ms.GetCollection(mid).FindOne(context.TODO(), bson.M{"_id": mid}).Decode(umsd)
	if err != mongo.ErrNoDocuments && err != nil {
		return umsd, err
	}
	ms.cache.Set(fmt.Sprint(mid), umsd, defaultExpire*time.Second)
	return umsd, nil
}

// Cache2DB 缓存数据回写
func (ms *Mstat) Cache2DB(ums *UserMstat) error {
	if ums.Mid == 0 {
		return errs.MidError
	}
	opts := options.Update().SetUpsert(true)
	_, err := ms.GetCollection(ums.Mid).UpdateOne(context.TODO(),
		bson.M{"_id": ums.Mid},
		bson.M{"$set": bson.M{
			"wtimes":    ums.Wtimes,
			"ltimes":    ums.Ltimes,
			"winmoney":  ums.WinMoney,
			"maxwin":    ums.MaxWin,
			"wlevel":    ums.Wlevel,
			"llevel":    ums.Llevel,
			"lgdays":    ums.Lgdays,
			"maxwin0":   ums.MaxWin0,
			"wmoney":    ums.Wmoney,
			"paytimes":  ums.PayTimes,
			"payamount": ums.PayAmount,
		}}, opts)
	ums.LastTime = time.Now()
	if err != nil {
		logger.Errorf("Cache2DB mid: %d, %v", ums.Mid, err)
		return err
	}
	return nil
}

// Close 关闭资源
func (ms *Mstat) Close() {
	err = ms.mongoClient.Disconnect(context.Background())
	logger.Infof("close mongo client %v", err)
	ms.cache.Flush()
	logger.Close()
}

// GetCollection 获取表对象
func (ms *Mstat) GetCollection(mid uint32) *mongo.Collection {
	return ms.mongoClient.Database(defaultDatabase).Collection(defaultCollection + strconv.Itoa(int(mid)%defaultCollectionNum))
}

// UserMstat 用户数据结构
type UserMstat struct {
	LastTime time.Time

	Mid       uint32         `bson:"_id"`
	Wtimes    map[string]int `bson:"wtimes"`    // 赢取的次数
	Ltimes    map[string]int `bson:"ltimes"`    // 输的次数
	WinMoney  map[string]int `bson:"winmoney"`  // 输赢的钱数
	MaxWin    map[string]int `bson:"maxwin"`    // 各游戏最大输取
	Wlevel    map[string]int `bson:"wlevel"`    // 场次赢的次数
	Llevel    map[string]int `bson:"llevel"`    // 场次输的次数
	Lgdays    int            `bson:"lgdays"`    // 登录的天数
	MaxWin0   int            `bson:"maxwin0"`   // 最大赢取
	MaxOwn    int            `bson:"maxown"`    // 最大拥有
	Wmoney    int            `bson:"wmoney"`    // 输赢的钱数
	PayTimes  int            `bson:"paytimes"`  // 支付次数
	PayAmount float64        `bson:"payamount"` // 支付金额
}

// MSet 批量设置数值
func (ums *UserMstat) MSet(value *map[string]interface{}) error {
	if len(*value) == 0 {
		return errs.EmptyError
	}

	ums.parseValue(value, util.ParseSet)
	return nil
}

// MIncrby 批量自增数值
func (ums *UserMstat) MIncrby(value *map[string]interface{}) error {
	if len(*value) == 0 {
		return errs.EmptyError
	}
	ums.parseValue(value, util.ParseIncrby)
	return nil
}

// MSetMax 批量最大值
func (ums *UserMstat) MSetMax(value *map[string]interface{}) error {
	if len(*value) == 0 {
		return errs.EmptyError
	}
	ums.parseValue(value, util.ParseSetMax)
	return nil
}

func (ums *UserMstat) parseValue(value *map[string]interface{}, parseType uint8) {
	for k, v := range *value {
		switch k {
		case "Wtimes":
			util.ParseValue(v, &ums.Wtimes, parseType)
		case "Ltimes":
			util.ParseValue(v, &ums.Ltimes, parseType)
		case "WinMoney":
			util.ParseValue(v, &ums.WinMoney, parseType)
		case "MaxWin":
			util.ParseValue(v, &ums.MaxWin, parseType)
		case "Wlevel":
			util.ParseValue(v, &ums.Wlevel, parseType)
		case "Llevel":
			util.ParseValue(v, &ums.Llevel, parseType)
		case "Lgdays":
			util.ParseValue(v, &ums.Lgdays, parseType)
		case "Wmoney":
			util.ParseValue(v, &ums.Wmoney, parseType)
		case "PayTimes":
			util.ParseValue(v, &ums.PayTimes, parseType)
		case "PayAmount":
			util.ParseValue(v, &ums.PayAmount, parseType)

		}
	}
}
