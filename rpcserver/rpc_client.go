package mylib

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func RpcClient() {
	client, err := rpc.DialHTTP("tcp", "192.168.56.101:1234")
	if err != nil {
		log.Fatal("dialing err", err)
	}

	args := &Args{16, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error ", err)
	}
	fmt.Printf("Arith: %d * %d = %d\n", args.A, args.B, reply)

	quotient := new(Quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	replyCall := <-divCall.Done

	if replyCall.Error != nil {
		log.Fatal("arith error ", replyCall.Error)
	}
	fmt.Printf("Arith: %d / %d = %d ... %d", args.A, args.B, quotient.Quo, quotient.Rem)
}
