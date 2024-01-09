package main

import (
	"database/sql"
	"net"
	db "service_payment/db/sqlc"
	"service_payment/gapi"
	"service_payment/pb"
	"service_payment/util/config"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Config struct {
	db *sql.DB
}

func runGrpcServer(config config.Config, store db.Store) {
	server := gapi.NewServer(config, store)

	grpcServer := grpc.NewServer()

	pb.RegisterPaymentServiceServer(grpcServer, server)

	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)

	if err != nil {
		panic(err)
	}

	errServer := grpcServer.Serve(listener)

	if errServer != nil {
		panic(errServer)
	}

}

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		panic(err)
	}

	cnn, err := connectToDB(config.DSN)

	if err != nil {
		panic(err)
	}

	store := db.NewStore(cnn)

	if err != nil {
		panic(err)

	}

	runGrpcServer(*config, *store)

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
