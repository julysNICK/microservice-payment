package gapi

import (
	"service_email/pb"
	"service_email/util/config"
)

type Server struct {
	pb.UnimplementedEmailServiceServer
	config config.Config
}

func NewServer(config config.Config) (*Server, error) {

	server := &Server{
		config: config,
	}

	return server, nil
}
