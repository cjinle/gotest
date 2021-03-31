package mstat

import (
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func TestOne(t *testing.T) {
	ms := New()
	mid := uint32(110)
	ums, err := ms.GetCache(mid)
	if err != nil {
		log.Fatalln(err)
	}
	val := make(map[string]interface{})
	val["PayTimes"] = 150
	val["PayAmount"] = float64(999.99)
	log.Println(ums.MSet(&val))
	val = map[string]interface{}{}
	val["Lgdays"] = 1
	log.Println(ums.MIncrby(&val))
	log.Println(ms.Cache2DB(ums))
}

func BenchmarkOne(b *testing.B) {
	ms := New()
	mid := uint32(110)

	for i := 0; i < b.N; i++ {
		ums, err := ms.GetCache(mid)
		if err != nil {
			log.Fatalln(err)
		}
		val := make(map[string]interface{})
		val["Lgdays"] = 1
		ums.MIncrby(&val)
		ms.Cache2DB(ums)
	}
}
