package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/cjinle/test/grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewFooClient(conn)

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// go func() {
	// 	select {
	// 	case <-ctx.Done():
	// 		log.Println("xxxxxxxxxxxxx")
	// 		os.Exit(0)
	// 	}
	// }()
	ctx := context.Background()
	var i int32

	go func() {
		for {
			i++
			r, err := c.Add(ctx, &pb.Args{Num1: int32(rand.Intn(100)), Num2: int32(rand.Intn(1000))})
			if err != nil {
				log.Println(err)
			}
			log.Println("reply: ", r.Num)
			time.Sleep(time.Second * 1)
		}
	}()

	cli, err := c.Pipe(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	go func() {
		i = 0
		for {
			i++
			err = cli.Send(&pb.DataPack{Cmd: i, Data: []byte(fmt.Sprintf("num: %d", rand.Int31()))})
			if err != nil {
				log.Println(err)
				return
			}
			time.Sleep(time.Second * 1)
		}
	}()
	go func() {
		for {
			r2, err := cli.Recv()
			if err != nil {
				log.Println(err)
				return
			}
			log.Println(r2.GetCmd(), string(r2.GetData()))

		}
	}()
	select {}

}
