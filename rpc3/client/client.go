package client

import (
	"log"
	"math/rand"
	"net/rpc"
	"time"
)

type Args struct {
	Mid uint32
}

type Reply struct {
	Ret  int32
	Data []byte
}

type User struct {
	Mid    uint32
	Mnick  string
	Gender uint8
	Sicon  string
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func RpcClient() {
	client, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().Unix())
	for {
		args := &Args{Mid: uint32(100000 + rand.Intn(10000))}
		log.Println("args: ", args)
		var reply Reply
		err = client.Call("User.GetInfo", args, &reply)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(args, reply, string(reply.Data))
		time.Sleep(1 * time.Second)
	}
}

func RpcGatewayClient() {

}
