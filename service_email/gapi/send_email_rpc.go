package gapi

import (
	"context"
	"fmt"
	"os"
	"service_email/cmd"
	"service_email/pb"
)

func (s *Server) SendEmail(ctx context.Context, req *pb.SendEmailRequest) (*pb.SendEmailResponse, error) {
	absPath, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("get current path error: %w", err)
	}

	attachFile := []string{absPath + "/README.md"}
	sender := cmd.NewGmailSender(s.config.EmailSenderName, s.config.EmailSenderAddress, s.config.EmailSenderPassword)

	toConvert := []string{req.To}

	err = sender.SenderEmail(req.Subject, req.Html, toConvert, nil, nil, attachFile)

	if err != nil {
		fmt.Println("send email error: ", err)
		return nil, fmt.Errorf("send email error: %w", err)
	}

	rsp := &pb.SendEmailResponse{
		Id:      "1",
		Status:  "OK",
		Message: "Send email success",
	}

	return rsp, nil
}
