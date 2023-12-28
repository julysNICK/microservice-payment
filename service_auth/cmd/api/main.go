package main

import (
	"fmt"
	"net"
	"service_auth/gapi"
	"service_auth/pb"
	"service_auth/util/config"

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

	pb.RegisterAuthServiceServer(grpcServer, server)

	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)

	if err != nil {
		fmt.Println("failed to listen: ", err)
	}

	errServer := grpcServer.Serve(listener)

	if errServer != nil {
		fmt.Println("failed to serve: ", err)
	}

}

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		fmt.Println("failed to load config: ", err)
	}

	runGrpcServer(*config)
	fmt.Println("starting grpc server...")
}
