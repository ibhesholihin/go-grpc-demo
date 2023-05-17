package main

import (
	"context"

	pb "github.com/ibhesholihin/go-grpc-demo/proto"
)

func (s *helloServer) SayHello(ctx context.Context, rqb *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
