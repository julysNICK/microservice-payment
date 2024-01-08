package gapi

import (
	db "service_auth/db/sqlc"
	"service_auth/pb"
	"service_auth/util/config"
	"service_auth/util/token"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
	config config.Config
	store  db.Store
	token  token.Maker
}

func NewServer(config config.Config, store db.Store, token token.Maker) (*Server, error) {

	server := &Server{
		config: config,
		store:  store,
		token:  token,
	}

	return server, nil
}
