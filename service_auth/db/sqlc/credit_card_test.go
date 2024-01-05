package db

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

var ErrNotFound = errors.New("not found")

func TestInsertCreditCardMock(t *testing.T) {
	db, mock, err := sqlmock.New()

	require.NoError(t, err)

	defer db.Close()

	mockDb := New(db)

	arg := CreateCreditCardParams{
		UserID: 1,

		Number: "1234567891234567",

		Balance: 1000,

		CreatedAt: time.Now(),

		UpdatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "user_id", "number", "balance", "created_at", "updated_at"}).
		AddRow(1, arg.UserID, arg.Number, arg.Balance, arg.CreatedAt, arg.UpdatedAt)

	mock.ExpectQuery("INSERT INTO credit_card").
		WithArgs(arg.UserID, arg.Number, arg.Balance, arg.CreatedAt, arg.UpdatedAt).
		WillReturnRows(rows)

	creditCard, err := mockDb.CreateCreditCard(context.Background(), arg)

	require.NoError(t, err)

	require.NotEmpty(t, creditCard)
}

func TestGetCreditCardByUserIdMock(t *testing.T) {
	creditCards := []CreditCard{
		{
			ID:        1,
			UserID:    1,
			Number:    "1234567891234567",
			Balance:   1000,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			UserID:    1,
			Number:    "1234567891234567",
			Balance:   1000,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        3,
			UserID:    1,
			Number:    "1234567891234567",
			Balance:   1000,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	db, mock, err := sqlmock.New()

	require.NoError(t, err)

	defer db.Close()

	mockDb := New(db)

	rows := sqlmock.NewRows([]string{"id", "user_id", "number", "balance", "created_at", "updated_at"}).AddRow(creditCards[0].ID, creditCards[0].UserID, creditCards[0].Number, creditCards[0].Balance, creditCards[0].CreatedAt, creditCards[0].UpdatedAt)

	mock.ExpectQuery("SELECT (.+) FROM credit_card WHERE user_id = (.+)").WithArgs(creditCards[0].UserID).WillReturnRows(rows)

	creditCardResult, err := mockDb.GetCreditCardByUserID(context.Background(), creditCards[0].UserID)

	require.NoError(t, err)

	require.NotEmpty(t, creditCardResult)
}

func TestGetCreditCardByNumberMockErrorNotFound(t *testing.T) {

	creditCards := []CreditCard{
		{
			ID:        1,
			UserID:    1,
			Number:    "1234567891234567",
			Balance:   1000,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			UserID:    1,
			Number:    "1234567891234567",
			Balance:   1000,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        3,
			UserID:    1,
			Number:    "1234567891234567",
			Balance:   1000,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	db, mock, err := sqlmock.New()

	require.NoError(t, err)

	defer db.Close()

	mockDb := New(db)

	mock.ExpectQuery("SELECT (.+) FROM credit_card WHERE user_id = (.+)").WithArgs(creditCards[0].UserID).WillReturnError(ErrNotFound)

	_, err = mockDb.GetCreditCardByUserID(context.Background(), creditCards[0].UserID)

	require.Error(t, err)

}
