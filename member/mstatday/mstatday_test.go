package mstatday

import (
	"log"
	"math/rand"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func TestOne(t *testing.T) {
	msd := New()
	umsd, err := msd.GetCache(uint32(110))
	if err != nil {
		log.Println(err)
	}
	umsd.PayAmount += 100.99
	log.Println(umsd)
	m := map[string]interface{}{
		"PayTimes": 10,
		"PlayTime": 1000,
	}
	umsd.MSet(&m)
	log.Println(umsd)
	m2 := map[string]interface{}{
		"PayTimes":  10,
		"PlayTime":  1000,
		"PayAmount": 2.99,
	}
	umsd.MIncrby(&m2)
	log.Println(umsd)
	// for {
	// 	m2 := map[string]interface{}{
	// 		"PayTimes":  10,
	// 		"PlayTime":  1000,
	// 		"PayAmount": 2.99,
	// 	}
	// 	umsd.MIncrby(&m2)
	// 	log.Println(umsd)
	// 	time.Sleep(time.Second)
	// }

}

func BenchmarkOne(b *testing.B) {
	msd := New()
	for i := 0; i < b.N; i++ {
		umsd, err := msd.GetCache(uint32(rand.Intn(9999)))
		if err != nil {
			log.Println(err)
		}
		umsd.PayAmount += 100.99
		// log.Println(umsd)
		m := map[string]interface{}{
			"paytimes": 10,
		}
		umsd.MSet(&m)
		// log.Println(umsd)
		m2 := map[string]interface{}{
			"PayTimes":  10,
			"PlayTime":  1000,
			"PayAmount": 2.99,
		}
		umsd.MIncrby(&m2)
		// log.Println(umsd)
	}

}

func BenchmarkTwo(b *testing.B) {

}
