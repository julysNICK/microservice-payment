package gapi

import (
	"context"
	"service_auth/pb"
)

func (s *Server) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	rsp := &pb.HelloResponse{
		Message: "Hello " + req.Name,
	}

	return rsp, nil
}
