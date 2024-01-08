package gapi

import (
	"context"
	"fmt"
	"service_auth/pb"
	"service_auth/util"
	"time"
)

func (s *Server) LoginUser(ctx context.Context, req *pb.AuthLoginRequest) (*pb.AuthLoginResponse, error) {
	userResult, err := s.store.GetUserByUsername(ctx, req.Username)

	if err != nil {
		return nil, err
	}

	checkPassword := util.CheckPassword(userResult.Password, req.Password)

	if !checkPassword {
		return nil, fmt.Errorf("invalid password")
	}

	token, err := s.token.CreateToken(userResult.Email, time.Minute*15)

	if err != nil {
		return nil, err
	}

	rsp := &pb.AuthLoginResponse{
		Token: token,
	}

	return rsp, nil

}
