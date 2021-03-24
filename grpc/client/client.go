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
	i := 0

	for {
		i++
		r, err := c.Add(ctx, &pb.Args{Num1: int32(rand.Intn(100)), Num2: int32(rand.Intn(1000))})
		if err != nil {
			log.Println(err)
		}
		log.Println("reply: ", r.Num)

		time.Sleep(time.Second)
		if i > 5 {
			continue
		}

		cli, err := c.SayHello(ctx)
		if err != nil {
			log.Println(err)
			continue
		}
		err = cli.Send(&pb.HelloRequest{S: fmt.Sprintf("num: %d", r.Num)})
		if err != nil {
			log.Println(err)
			continue
		}
		r2, err := cli.Recv()
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println(r2.GetS())

	}

}
