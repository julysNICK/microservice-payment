package gapi

import (
	"context"
	"service_payment/pb"
)

func (s *Server) HelloRpc(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello " + req.Name}, nil
}
