package main

import (
	"database/sql"
	"fmt"
	"net"
	db "service_auth/db/sqlc"
	"service_auth/gapi"
	"service_auth/pb"
	"service_auth/util/config"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Config struct {
	db *sql.DB
}

func runGrpcServer(config config.Config, store db.Store) {

	fmt.Println("starting grpc server...")

	server, err := gapi.NewServer(config, store)

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
	config, err := config.LoadConfig("../../")

	if err != nil {
		fmt.Println("failed to load config: ", err)
	}

	cnn, err := connectToDB(config.DSN)

	if err != nil {
		fmt.Println("failed to load config: ", err)
	}

	store := db.NewStore(cnn)

	runGrpcServer(*config, *store)
	fmt.Println("starting grpc server...")
}

func connectToDB(dns string) (*sql.DB, error) {

	db, err := openDBConnection(dns)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func openDBConnection(dns string) (*sql.DB, error) {

	conn, err := sql.Open("postgres", dns)

	if err != nil {
		return nil, err
	}

	return conn, nil
}
