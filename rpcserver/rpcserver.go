package rpcserver

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
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
	fmt.Println("func Multiply ... ")
	fmt.Println(args)
	fmt.Println(*reply)
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}

	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	fmt.Println("func Divide ... ")
	fmt.Println(args)
	fmt.Println(quo)
	return nil
}

func RpcServer() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error ", err)
	}
	http.Serve(listen, nil)
}
