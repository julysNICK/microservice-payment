package gapi

import (
	"context"
	"service_auth/pb"
	"service_auth/util/validations"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) RegisterUser(ctx context.Context, req *pb.AuthRegisterRequest) (*pb.AuthRegisterResponse, error) {

	violations := validate(req)

	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}

	rsp := &pb.AuthRegisterResponse{
		Message: "requisition name " + req.Username + " requisition password " + req.Password + " requisition email " + req.Email + " requisition cpf " + req.Cpf + " requisition credit number " + req.CreditCard,
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
