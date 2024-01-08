package gapi

import (
	"context"
	"service_auth/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) TokenAuth(ctx context.Context, req *pb.TokenRequest) (*pb.TokenResponse, error) {

	_, err := s.token.VerifyToken(req.Token)

	if err != nil {

		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	rsp := &pb.TokenResponse{
		Valid: true,
		Error: "",
	}

	return rsp, nil
}
