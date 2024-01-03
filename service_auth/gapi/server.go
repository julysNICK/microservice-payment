package gapi

import (
	db "service_auth/db/sqlc"
	"service_auth/pb"
	"service_auth/util/config"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
	config config.Config
	store  db.Store
}

func NewServer(config config.Config, store db.Store) (*Server, error) {

	server := &Server{
		config: config,
		store:  store,
	}

	return server, nil
}
