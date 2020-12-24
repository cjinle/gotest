package main

import (
	"log"
	"math/rand"
	"net/rpc"
	"net/url"
	"time"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	urlInfo, err := url.Parse("unix:///tmp/rpc2.sock")
	if err != nil {
		log.Fatal(err)
	}
	client, err := rpc.Dial(urlInfo.Scheme, urlInfo.Path)
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