package main

import (
	hello "github.com/sejin-P/envoy-grpc-transcoder-poc"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	grpcServer := hello.NewGrpcServer()
	go func() {
		lis, err := net.Listen("tcp", ":"+"6565")
		if err != nil {
			log.Fatal(err)
		}

		if err := grpcServer.Serve(lis); err != nil && err != grpc.ErrServerStopped {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(quit)

	<-quit

	log.Printf("Stopping golang gRPC server")
	grpcServer.GracefulStop()
}
