package gapi

import (
	"context"
	"database/sql"
	"fmt"
	"service_auth/creditService"
	db "service_auth/db/sqlc"
	"strconv"
	"time"

	"google.golang.org/grpc/status"
)

func isPossibleBuy(balance int32, amount int32) bool {
	result := balance - amount

	if result < 0 {
		return false
	}

	return true
}

func updateBalance(balance int32, amount int32) int32 {
	result := balance - amount

	return result
}

func convertStringToInt64(value string) (int32, error) {

	convertStringToInt64, err := strconv.ParseInt(value, 10, 64)

	if err != nil {
		return 0, err
	}

	convertStringToInt32 := int32(convertStringToInt64)

	return convertStringToInt32, nil

}

func convertInt32ToString(value int32) string {
	convertInt32ToString := strconv.Itoa(int(value))

	return convertInt32ToString
}

func (s *Server) GetCardByUserID(ctx context.Context, req *creditService.GetCardByUserIDRequest) (*creditService.GetCardByUserIDResponse, error) {

	fmt.Println("GetCardByUserID")

	if req.UserId == "" {
		return nil, status.Errorf(400, "user id is required")
	}

	convertStringToInt32, err := convertStringToInt64(req.UserId)

	if err != nil {
		return nil, status.Errorf(400, "user id is invalid")
	}

	cards, err := s.store.GetCreditCardByUserID(ctx, convertStringToInt32)

	if err != nil {
		return nil, status.Errorf(500, "error to get credit card by user id")
	}

	for _, card := range cards {
		rsp := &creditService.GetCardByUserIDResponse{
			Cards: []*creditService.Card{
				{
					Id:        convertInt32ToString(card.ID),
					Number:    card.Number,
					UserId:    convertInt32ToString(card.UserID),
					Balance:   convertInt32ToString(card.Balance),
					CreatedAt: card.CreatedAt.String(),
					UpdatedAt: card.UpdatedAt.String(),
				},
			},
		}

		return rsp, nil
	}

	return nil, status.Errorf(404, "card not found")
}
func (s *Server) UpdateCardBalance(ctx context.Context, req *creditService.UpdateCardBalanceRequest) (*creditService.UpdateCardBalanceResponse, error) {
	fmt.Println("UpdateCardBalance")

	if req.Id == "" {
		return nil, status.Errorf(400, "card id is required")
	}

	if req.Balance == "" {
		return nil, status.Errorf(400, "balance is required")
	}

	convertStringToInt32, err := convertStringToInt64(req.Id)

	if err != nil {
		return nil, status.Errorf(400, "card id is invalid")
	}

	convertStringToInt32Balance, err := convertStringToInt64(req.Balance)

	if err != nil {
		return nil, status.Errorf(400, "balance is invalid")
	}

	cardFound, err := s.store.GetCreditCardByID(ctx, convertStringToInt32)

	if err != nil {
		return nil, status.Errorf(500, "error to get credit card by id")
	}

	if !isPossibleBuy(cardFound[0].Balance, convertStringToInt32Balance) {
		return nil, status.Errorf(400, "balance is invalid")
	}

	errGet := s.store.UpdateCreditCardPartial(ctx, db.UpdateCreditCardPartialParams{
		ID: convertStringToInt32,
		Number: sql.NullString{
			String: cardFound[0].Number,
			Valid:  true,
		},
		Balance: sql.NullInt32{
			Int32: updateBalance(cardFound[0].Balance, convertStringToInt32Balance),
			Valid: true,
		},

		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})

	if errGet != nil {
		return nil, status.Errorf(500, "error to update credit card")
	}

	rsp := &creditService.UpdateCardBalanceResponse{
		Id:        req.Id,
		Balance:   req.Balance,
		Number:    cardFound[0].Number,
		UpdatedAt: time.Now().String(),
		CreatedAt: cardFound[0].CreatedAt.String(),
	}

	return rsp, nil
}
