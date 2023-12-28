package gapi

import (
	"service_auth/pb"
	"service_auth/util/config"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
	config config.Config
}

func NewServer(config config.Config) (*Server, error) {

	server := &Server{
		config: config,
	}

	return server, nil
}
