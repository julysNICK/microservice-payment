package gapi

import (
	db "service_payment/db/sqlc"
	"service_payment/pb"
	"service_payment/util/config"
)

type Server struct {
	pb.UnimplementedPaymentServiceServer

	config config.Config
	store  db.Store
}

func NewServer(config config.Config, store db.Store) *Server {
	return &Server{
		config: config,
		store:  store,
	}
}
