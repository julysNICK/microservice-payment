package gapi

import (
	"context"
	"fmt"
	db "service_auth/db/sqlc"
	"service_auth/pb"
	"service_auth/util"
	"service_auth/util/validations"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) RegisterUser(ctx context.Context, req *pb.AuthRegisterRequest) (*pb.AuthRegisterResponse, error) {

	violations := validate(req)

	cryptPassword := util.HashedPassword(req.Password)

	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}

	idUserSaved, err := s.store.CreateUserTx(ctx, db.UserTxParams{
		Email:      req.Email,
		Cpf:        req.Cpf,
		Password:   cryptPassword,
		CreditCard: req.CreditCard,
	})

	rsp := &pb.AuthRegisterResponse{
		Message: "requisition name " + string(idUserSaved) + " requisition password " + req.Password + " requisition email " + req.Email + " requisition cpf " + req.Cpf + " requisition credit number " + req.CreditCard,
	}
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func validate(req *pb.AuthRegisterRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	validateCpfBool := validations.IsValidCPF(req.Cpf)
	validateCreditNumberBool := validations.IsValidCreditCard(req.CreditCard)
	validateNameBool := validations.ValidateName(req.Username)

	if !validateCpfBool {

		violations = append(violations, &errdetails.BadRequest_FieldViolation{
			Field:       "cpf",
			Description: "cpf invalid",
		})
	}

	if !validateCreditNumberBool {
		fmt.Println("credit card invalid")
		violations = append(violations, &errdetails.BadRequest_FieldViolation{
			Field:       "credit_card",
			Description: "credit card invalid",
		})
	}

	if !validateNameBool {
		violations = append(violations, &errdetails.BadRequest_FieldViolation{
			Field:       "username",
			Description: "username invalid",
		})
	}

	return
}

func fieldViolation(field string, description string) *errdetails.BadRequest_FieldViolation {
	return &errdetails.BadRequest_FieldViolation{
		Field:       field,
		Description: description,
	}
}

func invalidArgumentError(violations []*errdetails.BadRequest_FieldViolation) error {
	badRequest := &errdetails.BadRequest{
		FieldViolations: violations,
	}
	statusInvalid := status.New(codes.InvalidArgument, "invalid argument")

	statusInvalid, err := statusInvalid.WithDetails(badRequest)

	if err != nil {
		return err
	}

	return statusInvalid.Err()
}
