package main

import (
	"context"
	"io"
	"log"
	"net"

	"github.com/cjinle/test/grpc/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedFooServer
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func (s *server) Add(ctx context.Context, in *pb.Args) (*pb.Reply, error) {
	log.Println(in.GetNum1(), in.GetNum2())
	return &pb.Reply{Num: in.Num1 + in.Num2}, nil
}

func (s *server) SayHello(stream pb.Foo_SayHelloServer) error {
	go func() {
		select {
		case <-stream.Context().Done():
			log.Println("done")
		}
	}()
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			log.Println("read done")
			return nil
		}
		if err != nil {
			log.Println("err", err)
			return nil
		}
		log.Println(in.GetS())
		err = stream.Send(&pb.HelloReply{S: "reply999"})
		if err != nil {
			return nil
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterFooServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
