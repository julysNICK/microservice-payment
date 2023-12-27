package gapi

import (
	"context"
	"fmt"
	"service_email/pb"
)

func (s *Server) SendEmail(ctx context.Context, req *pb.SendEmailRequest) (*pb.SendEmailResponse, error) {

	fmt.Println("SendEmail")

	fmt.Println("req: ", req)

	rsp := &pb.SendEmailResponse{
		Id:      "123",
		Status:  "OK",
		Message: "Email sent successfully",
	}

	return rsp, nil
}
