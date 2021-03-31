package gameinfo_test

import (
	"log"
	"math/rand"
	"testing"

	"github.com/cjinle/game/lib/errs"
	"github.com/cjinle/test/member/gameinfo"
)

func TestMoney(t *testing.T) {
	mid := uint32(1000000) + uint32(rand.Intn(100000))
	mode := uint32(100)
	gi := gameinfo.New()
	// log.Println(gi.mongo)
	ugi, err := gi.GetCache(mid)
	if err == errs.RecordNotExist {
		ugi, err = gi.CreateUserGameInfo(mid)
		if err != nil {
			t.Error(err)
		}
	} else if err != nil {
		t.Error(err)
	}

	ugi.AddMoney(mode, int32(888888))
	log.Println(ugi)

	ugi.AddMoney(uint32(200), int32(-100))
	ugi.AddExp(1000)
	log.Println(ugi)
	// time.Sleep(5 * time.Second)
	ugi.AddMoney(uint32(200), int32(-100))
	ugi.AddExp(1000)
	log.Println(ugi)
	// time.Sleep(5 * time.Second)
}

func BenchmarkMoney(b *testing.B) {

	// log.Println(gi)
	// b.N = 1000
	gi := gameinfo.New()
	for i := 0; i < b.N; i++ {
		mid := uint32(1000000) + uint32(rand.Intn(10000))
		mode := uint32(100)
		ugi, err := gi.GetCache(mid)
		if err == errs.RecordNotExist {
			ugi, err = gi.CreateUserGameInfo(mid)
			if err != nil {
				// log.Println(err)
				continue
			}
		} else if err != nil {
			continue
			// log.Println(err)
		}

		ugi.AddMoney(mode, 1)
		// ugi.AddMoney(uint32(200), -100)
		// ugi.AddExp(1000)
		// // log.Println(ugi)
		// ugi.AddMoney(uint32(200), 100)
		// ugi.AddExp(-100)
		// log.Println(ugi)
	}
}
