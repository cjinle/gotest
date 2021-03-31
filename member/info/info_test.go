package info_test

import (
	"log"
	"testing"
	"time"

	"github.com/cjinle/game/lib/errs"
	"github.com/cjinle/test/member/info"
)

func TestOne(t *testing.T) {
	log.Println("info test ... ")

	mid := uint32(1000001)
	info := info.New()

	ui, err := info.GetCache(mid)
	if err == errs.RecordNotExist {
		ui, err = info.CreateUserInfo(mid)
		if err != nil {
			t.Error(err)
		}
	} else if err != nil {
		t.Error(err)
	}

	log.Println("Mid:", ui.Mid)
	log.Println("Mnick:", ui.Mnick)
	log.Println("Sex:", ui.Sex)
	log.Println("Mstatus:", ui.Mstatus)
	log.Println("Mtime:", ui.Mtime)
	log.Println("Sicon:", ui.Sicon)
	log.Println("Bicon:", ui.Bicon)
	log.Println("Email:", ui.Email)

	ui.UpdateMnick("李四")
	ui.UpdateSex(1)
	ui.UpdateMstatus(1)
	ui.UpdateSicon("http://www.baidu.com/xxx_100x100.png")
	ui.UpdateBicon("http://www.baidu.com/xxx_200x200.png")

	ui, _ = info.GetCache(mid)
	log.Println(ui)

	time.Sleep(1 * time.Second)
}

func TestTwo(t *testing.T) { // go test -run TestTwo
	guid := "60235f39f2985611b2aad462"
	info := info.New()
	ui, err := info.GetByGUID(guid)
	log.Println(ui, err)

	log.Println(info.GetCache(ui.Mid))
}

func TestGetLastMid(t *testing.T) {
	info := info.New()
	mid := info.GetLastMid()
	log.Println(mid)
	newMid, err := info.GenMid()
	log.Println(newMid, err)
	info.CreateUserInfo(newMid)
}
