package gateway

import (
	"log"
	"net"
	"net/rpc"
)

type Args struct {
	Method string

	Mid uint32
}

type Reply struct {
	Ret  int32
	Data []byte
}

type Gateway struct {
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func (gw *Gateway) Send(args *Args, reply *Reply) error {
	return nil
}

func RpcServer() {
	rpcServer := rpc.NewServer()
	gw := new(Gateway)
	rpcServer.Register(gw)
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatal("listen error ", err)
	}
	go rpcServer.Accept(listen)
	select {}
}
