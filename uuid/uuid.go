package uuid

import (
	"fmt"
	"hash/fnv"
	"log"
	"math/rand"
	"time"
	"hash/crc32"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UUID struct {
	t time.Time
	r int64
}

func (u *UUID) String() string {
	v := fmt.Sprintf("%d-%d", u.t.UnixNano(), u.r)
	hash := fnv.New64a()
	hash.Write([]byte(v))
	return fmt.Sprintf("%d", hash.Sum64())
}

func NewUUID() *UUID {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &UUID{
		t: time.Now(),
		r: r.Int63(),
	}
}

func NewGoogleUUID() {
	id := uuid.New()
	log.Println(id)
}

func NewObjectID() uint64 {
	objectID := primitive.NewObjectID().Hex()
	fmt.Println(objectID)
	id32 := crc32.ChecksumIEEE([]byte(objectID))
	fmt.Println(id32)
	id64 := uint64(1e14)
	now := time.Now()
	id64 += uint64(id32)
	id64 += uint64(now.YearDay()) * 1e10
	fmt.Println(now.YearDay())
	id64 += uint64(now.Year()%100) * 1e10
	fmt.Println(now.Year())
	return id64
}
