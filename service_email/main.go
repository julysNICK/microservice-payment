package main

import (
	"fmt"
	"net"
	"service_email/gapi"
	"service_email/pb"
	"service_email/util/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func runGrpcServer(config config.Config) {

	fmt.Println("starting grpc server...")

	server, err := gapi.NewServer(config)

	if err != nil {
		fmt.Println("failed to create server: ", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterEmailServiceServer(grpcServer, server)

	fmt.Println("starting grpc server... line 28")
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	fmt.Println("starting grpc server... line 32")
	if err != nil {
		fmt.Println("failed to listen: ", err)
	}

	fmt.Println("starting grpc server... line 37")

	errServer := grpcServer.Serve(listener)
	fmt.Println("starting grpc server... line 40")
	if errServer != nil {
		fmt.Println("failed to serve: ", err)
	}

	fmt.Println("starting grpc server... line 45")

	fmt.Println("starting grpc server...")
}

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		fmt.Println("failed to load config: ", err)
	}

	runGrpcServer(*config)
	fmt.Println("starting grpc server...")
}
