package mstat2

import (
	"log"
	"math/rand"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func TestOne(t *testing.T) {
	m := map[string]int{
		"lgdays":    10,
		"maxwin":    10000,
		"maxwin_1":  1000,
		"payAmount": 1000,
	}
	mid := uint32(110)
	ms := New()
	err = ms.MsetInt(mid, &m)
	if err != nil {
		log.Println(err)
	}

	log.Println(ms.MgetInt(mid, []string{"lgdays", "maxwin"}))

	m2 := map[string]float64{
		"payAmount": 100.9,
		"score":     100.1,
	}
	err = ms.MsetFloat(mid, &m2)
	if err != nil {
		log.Println(err)
	}
	log.Println(ms.MgetFloat(mid, []string{"payAmount", "score"}))
}

func BenchmarkOne(b *testing.B) {
	ms := New()
	for i := 0; i < b.N; i++ {
		m := map[string]int{
			"lgdays":    10,
			"maxwin":    10000,
			"maxwin_1":  1000,
			"payAmount": 1000,
		}
		mid := uint32(10000 + rand.Intn(1000))

		err = ms.MsetInt(mid, &m)
		if err != nil {
			log.Println(err)
		}
		ms.MgetInt(mid, []string{"lgdays", "maxwin"})
		// log.Println(ms.MgetInt(mid, []string{"lgdays", "maxwin"}))

		m2 := map[string]float64{
			"payAmount": 100.9,
			"score":     100.1,
		}
		err = ms.MsetFloat(mid, &m2)
		if err != nil {
			log.Println(err)
		}
		ms.MgetFloat(mid, []string{"payAmount", "score"})
		// log.Println(ms.MgetFloat(mid, []string{"payAmount", "score"}))
	}
}
