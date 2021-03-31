package recent

import (
	"log"
	"math/rand"
	"testing"
)

func TestOne(t *testing.T) {
	mid := uint32(100000 + rand.Intn(10000))
	recent := New()
	userRecent, err := recent.Get(mid)
	if err != nil {
		log.Println(err)
	}
	log.Println(userRecent)
	userRecent.API = 0x0100B200
	userRecent.Ver = "1.0.2"
	userRecent.Lang = 2
	userRecent.Mac = "00:00:00:00:00"
	userRecent.IP = "8.8.8.8"
	userRecent.Imei = "02000000000"
	userRecent.DeviceID = "xxx-xxx-xxx-xxx"
	userRecent.AndroidID = "000-000-000-000"
	recent.AddCache2DBChan(mid)
	// time.Sleep(5 * time.Second)
	// recent.Close()
}

func BenchmarkOne(b *testing.B) {
	recent := New()
	for i := 1; i < b.N; i++ {
		mid := uint32(100000 + rand.Intn(10000))
		userRecent, err := recent.Get(mid)
		if err != nil {
			log.Println(err)
		}
		// log.Println(userRecent)
		userRecent.API = 0x0100B200
		userRecent.Ver = "1.0.2"
		userRecent.Lang = 2
		userRecent.Mac = "00:00:00:00:00"
		userRecent.IP = "8.8.8.8"
		userRecent.Imei = "02000000000"
		userRecent.DeviceID = "xxx-xxx-xxx-xxx"
		userRecent.AndroidID = "000-000-000-000"
		recent.AddCache2DBChan(mid)
	}
}
