package info

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/cjinle/game/lib/errs"
	"github.com/patrickmn/go-cache"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	defaultExpire       = 3600
	defaultCleanInetval = 60
	defaultDatabase     = "member"
	defaultCollection   = "info"
)

// Info 数据信息结构
type Info struct {
	cache     *cache.Cache
	mongo     *mongo.Client
	mongoColl *mongo.Collection
	lastMid   uint32
}

var cache2DBChan chan uint32
var err error

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

// New 初始化一个Info对象
func New() *Info {
	info := &Info{}
	info.cache = cache.New(defaultExpire*time.Second, defaultCleanInetval*time.Second)
	info.mongo, err = mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatalln("mongo connect err", err)
	}
	info.mongoColl = info.mongo.Database(defaultDatabase).Collection(defaultCollection)
	cache2DBChan = make(chan uint32, 1000)
	go func(info *Info) {
		for {
			select {
			case mid := <-cache2DBChan:
				err := info.Cache2DB(mid)
				log.Println(">>>> cache2DBChan", mid, err)
			}
		}
	}(info)
	return info
}

// GetCache 获取用户缓存数据
func (info *Info) GetCache(mid uint32) (*UserInfo, error) {
	cacheKey := fmt.Sprint(mid)
	if data, ok := info.cache.Get(cacheKey); ok {
		return data.(*UserInfo), nil
	}

	// load db to cache
	ui := &UserInfo{}
	err := info.mongoColl.FindOne(context.TODO(), bson.M{"mid": mid}).Decode(ui)
	if err != nil {
		return &UserInfo{}, errs.RecordNotExist
	}

	err = info.SetCache(mid, ui)
	if err != nil {
		log.Println(err)
	}
	return ui, nil
}

// SetCache 更新用户信息到缓存
func (info *Info) SetCache(mid uint32, ui *UserInfo) error {
	if mid != ui.Mid {
		return fmt.Errorf("mid: %d set cache error", mid)
	}
	info.cache.Set(fmt.Sprint(mid), ui, defaultExpire*time.Second)
	return nil
}

// GetByGUID 通过guid查找用户信息
func (info *Info) GetByGUID(guid string) (ui *UserInfo, err error) {
	if len(guid) == 0 {
		err = errors.New("guid empty")
		return
	}
	oid, err := primitive.ObjectIDFromHex(guid)
	if err != nil {
		return
	}
	ui = &UserInfo{}
	err = info.mongoColl.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(ui)
	if err != nil {
		return
	}
	cacheKey := fmt.Sprint(ui.Mid)
	if _, ok := info.cache.Get(cacheKey); !ok {
		info.SetCache(ui.Mid, ui)
	}
	return
}

// CreateUserInfo 创建用户
func (info *Info) CreateUserInfo(mid uint32) (*UserInfo, error) {
	ui := newUserInfo(mid)
	_, err := info.mongoColl.InsertOne(context.TODO(), bson.M{
		"_id":   primitive.NewObjectID(),
		"mid":   ui.Mid,
		"mtime": ui.Mtime,
	})
	return ui, err
}

// Cache2DB 缓存数据回到DB
func (info *Info) Cache2DB(mid uint32) error {
	cacheKey := fmt.Sprint(mid)
	if data, ok := info.cache.Get(cacheKey); ok {
		ui := data.(*UserInfo)
		_, err := info.mongoColl.UpdateOne(context.TODO(),
			bson.M{"mid": mid},
			bson.M{"$set": bson.M{
				"mnick":   ui.Mnick,
				"sex":     ui.Sex,
				"mstatus": ui.Mstatus,
				"mtime":   ui.Mtime,
				"sicon":   ui.Sicon,
				"bicon":   ui.Bicon,
				"email":   ui.Email,
			}})
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("Mid: %d not in cache", mid)
	}
	return nil
}

// GetLastMid 获取当前最大的mid
func (info *Info) GetLastMid() uint32 {
	if info.lastMid > 0 {
		return info.lastMid
	}
	queryOptions := options.FindOneOptions{}
	queryOptions.SetSort(bson.M{"_id": -1})
	ui := &UserInfo{}
	err = info.mongoColl.FindOne(context.TODO(), bson.D{}, &queryOptions).Decode(ui)
	if err != nil {
		return 0
	}
	info.lastMid = ui.Mid
	return info.lastMid
}

func (info *Info) GenMid() (mid uint32, err error) {
	maxMid := info.GetLastMid()
	if maxMid == 0 {
		err = errors.New("max mid error")
		return
	}
	info.lastMid++
	mid = info.lastMid
	return
}

// UserInfo 用户基本数据结构
type UserInfo struct {
	GUID    string `bson:"_id"`
	Mid     uint32 `bson:"mid"`
	Mnick   string `bson:"mnick"`
	Sex     uint8  `bson:"sex"`
	Mstatus uint8  `bson:"mstatus"`
	Mtime   uint32 `bson:"mtime"`
	Sicon   string `bson:"sicon"`
	Bicon   string `bson:"bicon"`
	Email   string `bson:"email"`
	Mu      sync.Mutex
	InChan  bool
}

func newUserInfo(mid uint32) *UserInfo {
	return &UserInfo{Mid: mid, Mtime: uint32(time.Now().Unix())}
}

// Update 更新用户信息
func (ui *UserInfo) Update() error {
	ui.Mu.Lock()
	defer ui.Mu.Unlock()
	ui.AddChan()
	return nil
}

// UpdateMnick 更新头像
func (ui *UserInfo) UpdateMnick(mnick string) error {
	if mnick == "" {
		return fmt.Errorf("Mid: %d empty mnick", ui.Mid)
	}
	ui.Mu.Lock()
	defer ui.Mu.Unlock()
	ui.Mnick = mnick
	ui.AddChan()
	return nil
}

// UpdateSex 更新性别
func (ui *UserInfo) UpdateSex(sex uint8) error {
	ui.Mu.Lock()
	defer ui.Mu.Unlock()
	ui.Sex = sex
	ui.AddChan()
	return nil
}

// UpdateMstatus 更新用户状态，0正常，1封号
func (ui *UserInfo) UpdateMstatus(mstatus uint8) error {
	ui.Mu.Lock()
	defer ui.Mu.Unlock()
	ui.Mstatus = mstatus
	ui.AddChan()
	return nil
}

// UpdateSicon 更新小头像
func (ui *UserInfo) UpdateSicon(icon string) error {
	ui.Mu.Lock()
	defer ui.Mu.Unlock()
	ui.Sicon = icon
	ui.AddChan()
	return nil
}

// UpdateBicon 更新大头像
func (ui *UserInfo) UpdateBicon(icon string) error {
	ui.Mu.Lock()
	defer ui.Mu.Unlock()
	ui.Bicon = icon
	ui.AddChan()
	return nil
}

// UpdateEmail 更新邮箱
func (ui *UserInfo) UpdateEmail(email string) error {
	ui.Mu.Lock()
	defer ui.Mu.Unlock()
	ui.Email = email
	ui.AddChan()
	return nil
}

// AddChan 加到缓存队列
func (ui *UserInfo) AddChan() {
	if !ui.InChan {
		ui.InChan = true
		cache2DBChan <- ui.Mid
	}
}
