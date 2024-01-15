package gapi

import (
	"context"
	"fmt"

	"service_payment/creditService"
	db "service_payment/db/sqlc"
	"service_payment/pb"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func (s *Server) PaymentCreate(ctx context.Context, req *pb.PaymentCreateRequest) (*pb.PaymentCreateResponse, error) {

	if req.Amount == "" {
		return nil, status.Errorf(400, "amount is required")
	}

	if req.StoreName == "" {
		return nil, status.Errorf(400, "store name is required")
	}

	if req.ProductId == "" {
		return nil, status.Errorf(400, "user id is required")
	}

	if req.UserId == "" {
		return nil, status.Errorf(400, "user id is required")
	}

	conn, err := grpc.Dial("localhost:9091", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, status.Errorf(500, "could not connect to user service grpc.Dial: %v", err)
	}

	defer conn.Close()

	reqCard := &creditService.GetCardByUserIDRequest{
		UserId: req.GetUserId(),
	}

	c := creditService.NewCreditServiceClient(conn)

	res, err := c.GetCardByUserID(ctx, reqCard)

	if err != nil {
		return nil, status.Errorf(500, "could not get card by user id GetCardByUserID: %v", err)
	}

	if res.Cards[0].Balance < req.Amount {

		s.store.CreatePurchase(ctx, db.CreatePurchaseParams{
			UserID:    convertStringToInt(req.UserId),
			ProductID: convertStringToInt(req.ProductId),
			StoreName: req.StoreName,
			Amount:    req.Amount,
			Status:    "failed: insufficient funds",
		})

		return nil, status.Errorf(400, "insufficient funds")
	}

	if res.Cards[0].UserId != req.UserId {

		s.store.CreatePurchase(ctx, db.CreatePurchaseParams{
			UserID:    convertStringToInt(req.UserId),
			ProductID: convertStringToInt(req.ProductId),
			StoreName: req.StoreName,
			Amount:    req.Amount,
			Status:    fmt.Sprintf("failed: user id does not match %v", res.Cards[0].UserId),
		})
		return nil, status.Errorf(400, "user id does not match")
	}

	_, err = s.store.CreatePurchase(ctx, db.CreatePurchaseParams{
		UserID:    convertStringToInt(req.UserId),
		ProductID: convertStringToInt(req.ProductId),
		StoreName: req.StoreName,
		Amount:    req.Amount,
		Status:    "success",
	})

	if err != nil {
		return nil, status.Errorf(500, "could not create purchase CreatePurchase: %v", err)
	}

	_, err = c.UpdateCardBalance(ctx, &creditService.UpdateCardBalanceRequest{
		Id:      res.Cards[0].Id,
		Balance: req.Amount,
	})

	if err != nil {
		s.store.CreatePurchase(ctx, db.CreatePurchaseParams{
			UserID:    convertStringToInt(req.UserId),
			ProductID: convertStringToInt(req.ProductId),
			StoreName: req.StoreName,
			Amount:    req.Amount,
			Status:    fmt.Sprintf("failed: %v", err),
		})
		return nil, status.Errorf(500, "could not update card balance UpdateCardBalance: %v", err)
	}

	return &pb.PaymentCreateResponse{
		Message: "success",
	}, nil
}

func convertStringToInt(str string) int32 {
	converted, _ := strconv.Atoi(str)

	return int32(converted)
}
