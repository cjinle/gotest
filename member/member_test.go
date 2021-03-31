package member_test

import (
	"log"
	"math/rand"
	"net/rpc"
	"testing"
	"time"

	"github.com/cjinle/test/member/rpcdata"
)

func TestOne(t *testing.T) {

	// urlInfo, err := url.Parse("unix:///tmp/rpc2.sock")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// client, err := rpc.Dial(urlInfo.Scheme, urlInfo.Path)
	client, err := rpc.Dial("tcp", "192.168.56.101:8888")
	if err != nil {
		log.Fatal(err)
	}
	rand.Seed(time.Now().Unix())

	for {
		mid := uint32(1000000)
		mid += uint32(rand.Intn(1000))
		param := &rpcdata.Args{
			Mid: mid,
		}
		log.Println(">>>> 1 args: ", param)
		var reply rpcdata.GameInfo
		err = client.Call("Member.GetGameInfo", param, &reply)
		if err != nil {
			log.Println(err)
		}
		log.Printf("<<<< Mid: %d, Money: %d\n", param.Mid, reply.Money)

		// param2 := &rpcdata.ArgsAddMoney{
		// 	Mid:  mid,
		// 	Mode: uint32(100),
		// 	Num:  int32(rand.Intn(1000)),
		// }
		// log.Println(">>>> 2 args2: ", param)
		// var reply2 rpcdata.ReplyAddMoney
		// err = client.Call("Member.AddMoney", param2, &reply2)
		// if err != nil {
		// 	log.Println(err)
		// }
		// log.Printf("<<<< Mid: %d, Money: %d\n", reply2.Mid, reply2.Money)
		// time.Sleep(1 * time.Second)
	}

}
