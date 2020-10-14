package recallstat

import (
	"fmt"
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	fmt.Println(ReadFile("recall_push.log"))
}

func TestDecode(t *testing.T) {
	str := `{"multicast_id":5138746323374678016,"success":2,"failure":5,"canonical_ids":0,"results":[{"error":"NotRegistered"},{"message_id":"0:1601379903326124%eb2476bdf9fd7ecd"},{"message_id":"0:1601379903326179%eb2476bdf9fd7ecd"},{"error":"NotRegistered"},{"error":"NotRegistered"},{"error":"NotRegistered"},{"error":"NotRegistered"}]}`
	msg, err := Decode(str)
	fmt.Printf("%+v", msg)
	fmt.Println(msg.MulticastId, err)
}

func TestTime(t *testing.T) {
	// fmt.Println(time.Now().Year())
	y, m, d := time.Now().Year(), time.Now().Month(), time.Now().Day()
	fmt.Println(y, m, d)
	fmt.Println(time.Now().Unix())
}
