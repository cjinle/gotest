package server

import (
	"encoding/json"
	"errors"
	"log"
	"net"
	"net/rpc"
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

func (user *User) GetInfo(args *Args, reply *Reply) error {
	// realArgs, ok := args.(Args)
	realArgs := args
	ok := true
	if !ok {
		log.Println(args)
		reply.Ret = -100
		return errors.New("args error")
	}
	user.Mid = realArgs.Mid
	user.Mnick = "hello"
	user.Gender = 1
	user.Sicon = "http://www.baidu.com/xx.png"
	var err error
	reply.Data, err = json.Marshal(user)
	return err
}

func RpcServer() {
	rpcServer := rpc.NewServer()
	user := new(User)
	rpcServer.Register(user)
	listen, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Fatal("listen error ", err)
	}
	go rpcServer.Accept(listen)
	select {}
}
