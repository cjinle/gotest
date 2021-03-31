package mstatday

import (
	"context"
	"fmt"
	"log"
	"os"
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
	defaultExpire       = 3600
	defaultCleanInetval = 60
	defaultDatabase     = "mstatday"
	defaultCollection   = "mstatday"
	expire              = 7 // 缓存过期时间
)

// MstatDay 每日用户数据统计结构
type MstatDay struct {
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
			Prefix: "mstatday",
		},
	)
}

// New 初始化对象
func New() *MstatDay {
	msd := &MstatDay{}
	msd.cache = cache.New(defaultExpire*time.Second, defaultCleanInetval*time.Second)
	msd.cache.OnEvicted(func(key string, v interface{}) {
		err = msd.Cache2DB(v.(*UserMstatDay))
		if err != nil {
			logger.Errorf("cache.OnEvicted error mid: %s, %v %v", key, v, err)
		}
	})
	msd.mongoClient, err = mongo.Connect(context.Background(),
		options.Client().ApplyURI("mongodb://127.0.0.1:27017/?maxPoolSize=10"))
	if err != nil {
		logger.Fatalf("%v", err)
	}
	cache2DBChan = make(chan uint32, 1000)
	go func(msd *MstatDay) {
		for {
			select {
			case mid := <-cache2DBChan:
				cacheKey := fmt.Sprint(mid)
				if data, ok := msd.cache.Get(cacheKey); ok {
					err = msd.Cache2DB(data.(*UserMstatDay))
					if err != nil {
						logger.Errorf("cache2DBChan error mid: %d, %v %v", mid, data, err)
					}
				}
			}
		}
	}(msd)
	return msd
}

// GetCache 获取用户数据对象
func (msd *MstatDay) GetCache(mid uint32) (*UserMstatDay, error) {
	cacheKey := fmt.Sprint(mid)
	if data, ok := msd.cache.Get(cacheKey); ok {
		return data.(*UserMstatDay), nil
	}
	umsd := &UserMstatDay{Mid: mid}
	date := fmt.Sprintf("%d%d%d", time.Now().Year(), time.Now().Month(), time.Now().Day())
	err := msd.GetCollection(date).FindOne(context.TODO(), bson.M{"mid": mid}).Decode(umsd)
	if err != mongo.ErrNoDocuments && err != nil {
		return umsd, err
	}
	msd.cache.Set(fmt.Sprint(mid), umsd, defaultExpire*time.Second)
	return umsd, nil
}

// Cache2DB 缓存数据回写
func (msd *MstatDay) Cache2DB(umsd *UserMstatDay) error {
	log.Println(umsd, umsd.Mid)
	if umsd.Mid == 0 {
		return errs.MidError
	}
	opts := options.Update().SetUpsert(true)
	_, err := msd.GetCollection("2020").UpdateOne(context.TODO(),
		bson.M{"mid": umsd.Mid},
		bson.M{"$set": bson.M{}}, opts)
	umsd.LastTime = time.Now()
	if err != nil {
		logger.Errorf("Cache2DB mid: %d, %v", umsd.Mid, err)
		return err
	}
	return nil
}

// Close 关闭资源
func (msd *MstatDay) Close() {
	err = msd.mongoClient.Disconnect(context.Background())
	logger.Infof("close mongo client %v", err)
	msd.cache.Flush()
	logger.Close()
}

// GetCollection 获取表对象
func (msd *MstatDay) GetCollection(date string) *mongo.Collection {
	return msd.mongoClient.Database(defaultDatabase).Collection(defaultCollection + date)
}

// UserMstatDay 用户数据结构
type UserMstatDay struct {
	LastTime time.Time

	Mid             uint32
	Wtimes          map[string]int // 赢取的次数
	Ltimes          map[string]int // 输的次数
	WinMoney        map[string]int // 输赢的钱数
	MaxWin          map[string]int // 各游戏最大输取
	Wlevel          map[string]int // 场次赢的次数
	Llevel          map[string]int // 场次输的次数
	Wmoney          int            // 输赢的钱数
	PayTimes        int            // 支付次数
	PayAmount       float64        // 支付金额
	BoxTimes        int            // 当前记时宝箱的次数
	BoxTime         int            // 上次记时宝箱领取的时间
	BoxVideoTimes   int            // 看广告减少记时宝箱次数
	SendInvite      int            // 成功发送邀请人数
	PlayTime        int            // 在玩时长
	Invite          int            // 成功邀请人数
	RecallNum       int            // 召回了多少人
	SendRecall      int            // 发送召回人数
	SlotsTimes      int            // 玩老虎机次数
	SlotsWin        int            // 老虎机输赢
	WheelTimes      int            // 大转盘免费次数
	PlayWheel       int            // 玩大转盘次数
	CashcowSelf     int            // 摇钱树收取自己筹码
	CashcowFriend   int            // 摇钱树收取好友筹码
	ActIntegral     float64        // 积分活动
	InviteRewardNum int            // 邀请码奖励次数
	InviteSuccNum   int            // 邀请成功次数
	AdmobNum        int            // admob广告看视频次数
	TimePromoShow   int            // 限时特惠当天首次展示时间
	Bankrupt        int            // 破产次数
}

// MSet 批量设置数值
func (umsd *UserMstatDay) MSet(value *map[string]interface{}) error {
	if len(*value) == 0 {
		return errs.EmptyError
	}

	umsd.parseValue(value, util.ParseSet)
	return nil
}

// MIncrby 批量自增数值
func (umsd *UserMstatDay) MIncrby(value *map[string]interface{}) error {
	if len(*value) == 0 {
		return errs.EmptyError
	}
	umsd.parseValue(value, util.ParseIncrby)
	return nil
}

// MSetMax 批量最大值
func (umsd *UserMstatDay) MSetMax(value *map[string]interface{}) error {
	if len(*value) == 0 {
		return errs.EmptyError
	}
	umsd.parseValue(value, util.ParseSetMax)
	return nil
}

func (umsd *UserMstatDay) parseValue(value *map[string]interface{}, parseType uint8) {
	for k, v := range *value {
		switch k {
		case "Wtimes":
			util.ParseValue(v, &umsd.Wtimes, parseType)
		case "Ltimes":
			util.ParseValue(v, &umsd.Ltimes, parseType)
		case "WinMoney":
			util.ParseValue(v, &umsd.WinMoney, parseType)
		case "MaxWin":
			util.ParseValue(v, &umsd.MaxWin, parseType)
		case "Llevel":
			util.ParseValue(v, &umsd.Llevel, parseType)
		case "Wmoney":
			util.ParseValue(v, &umsd.Wmoney, parseType)
		case "PayTimes":
			util.ParseValue(v, &umsd.PayTimes, parseType)
		case "PayAmount":
			util.ParseValue(v, &umsd.PayAmount, parseType)
		case "BoxTimes":
			util.ParseValue(v, &umsd.BoxTimes, parseType)
		case "BoxTime":
			util.ParseValue(v, &umsd.BoxTime, parseType)
		case "BoxVideoTimes":
			util.ParseValue(v, &umsd.BoxVideoTimes, parseType)
		case "SendInvite":
			util.ParseValue(v, &umsd.SendInvite, parseType)
		case "PlayTime":
			util.ParseValue(v, &umsd.PlayTime, parseType)
		case "Invite":
			util.ParseValue(v, &umsd.Invite, parseType)
		case "RecallNum":
			util.ParseValue(v, &umsd.RecallNum, parseType)
		case "SendRecall":
			util.ParseValue(v, &umsd.SendRecall, parseType)
		case "SlotsTimes":
			util.ParseValue(v, &umsd.SlotsTimes, parseType)
		case "SlotsWin":
			util.ParseValue(v, &umsd.SlotsWin, parseType)
		case "WheelTimes":
			util.ParseValue(v, &umsd.WheelTimes, parseType)
		case "PlayWheel":
			util.ParseValue(v, &umsd.PlayWheel, parseType)
		case "CashcowSelf":
			util.ParseValue(v, &umsd.CashcowSelf, parseType)
		case "CashcowFriend":
			util.ParseValue(v, &umsd.CashcowFriend, parseType)
		case "ActIntegral":
			util.ParseValue(v, &umsd.ActIntegral, parseType)
		case "InviteRewardNum":
			util.ParseValue(v, &umsd.InviteRewardNum, parseType)
		case "InviteSuccNum":
			util.ParseValue(v, &umsd.InviteSuccNum, parseType)
		case "AdmobNum":
			util.ParseValue(v, &umsd.AdmobNum, parseType)
		case "TimePromoShow":
			util.ParseValue(v, &umsd.TimePromoShow, parseType)
		case "Bankrupt":
			util.ParseValue(v, &umsd.Bankrupt, parseType)
		}
	}
}
