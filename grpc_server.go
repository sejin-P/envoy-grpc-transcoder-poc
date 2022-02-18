package hello

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type helloServer struct {
	GreeterServer
}

func NewHelloServer() *helloServer {
	return &helloServer{}
}

func (s *helloServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	fmt.Println("requested")
	fmt.Println(req)
	if req.Hey.GetMe() != nil {
		return &HelloReply{
			First:  "me",
			Second: "me",
		}, nil
	}

	if req.Hey.GetYou() != nil {
		return &HelloReply{
			First:  "you",
			Second: "you",
		}, nil
	}

	return nil, status.Error(codes.Internal, "abc")
}

func NewGrpcServer() *grpc.Server {
	grpcServer := grpc.NewServer()

	helloServer := NewHelloServer()
	RegisterGreeterServer(grpcServer, helloServer)
	reflection.Register(grpcServer)

	return grpcServer
}
