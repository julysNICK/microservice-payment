package main

import (
	"fmt"
	"service_email/gapi"
	"service_email/pb"
	"service_email/util/config"

	"google.golang.org/grpc"
)

func runGrpcServer(config config.Config) {

	server, err := gapi.NewServer(config)

	if err != nil {
		fmt.Println("failed to create server: ", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterEmailServiceServer(grpcServer, server)

	fmt.Println("starting grpc server...")
}

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		fmt.Errorf("failed to load config: %w", err)
	}

	runGrpcServer(*config)
}
