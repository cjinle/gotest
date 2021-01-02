package main

import (
	"context"
	"log"
	"math/rand"
	"os"
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		select {
		case <-ctx.Done():
			log.Println("xxxxxxxxxxxxx")
			os.Exit(0)
		}
	}()

	for {
		r, err := c.Add(ctx, &pb.Args{Num1: int32(rand.Intn(100)), Num2: int32(rand.Intn(1000))})
		if err != nil {
			log.Println(err)
		}
		log.Println("reply: ", r.Num)
		time.Sleep(time.Second)
	}

}
