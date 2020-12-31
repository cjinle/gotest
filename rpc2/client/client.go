package main

import (
	"log"
	"math/rand"
	"net/rpc"
	"time"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	// urlInfo, err := url.Parse("unix:///tmp/gateway.sock")

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// client, err := rpc.Dial(urlInfo.Scheme, urlInfo.Path)
	// urlInfo, err := url.Parse("tcp://127.0.0.1:8888")
	// log.Println(urlInfo.Host, urlInfo.Scheme, err)
	// os.Exit(0)
	client, err := rpc.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().Unix())

	for {
		args := &Args{rand.Int(), rand.Int()}
		log.Println("args: ", args)
		var reply int
		err = client.Call("Arith.Multiply", args, &reply)
		if err != nil {
			log.Fatal("arith error ", err)
		}
		log.Printf("Arith: %d * %d = %d\n", args.A, args.B, reply)

		quotient := new(Quotient)
		divCall := client.Go("Arith.Divide", args, quotient, nil)
		replyCall := <-divCall.Done

		if replyCall.Error != nil {
			log.Fatal("arith error ", replyCall.Error)
		}
		log.Printf("Arith: %d / %d = %d ... %d", args.A, args.B, quotient.Quo, quotient.Rem)
		time.Sleep(1 * time.Second)
	}

}
