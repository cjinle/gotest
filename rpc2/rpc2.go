package rpc2

import (
	"errors"
	"log"
	"net"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}

	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func RpcServer() {
	rpcServer := rpc.NewServer()
	arith := new(Arith)
	rpcServer.Register(arith)
	// listen, err := net.Listen("unix", "/tmp/rpc2.sock")
	listen, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Fatal("listen error ", err)
	}
	go rpcServer.Accept(listen)
	select {}
}
