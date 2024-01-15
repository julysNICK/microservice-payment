package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func initTestDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}

func TestCreatePayment(t *testing.T) {
	db, mock := initTestDB(t)

	defer db.Close()

	rows := sqlmock.NewRows([]string{"purchase_id", "user_id", "product_id", "store_name", "amount", "status", "purchase_time", "created_at", "updated_at"}).
		AddRow(1, 1, 1, "Test Store", "10000", "pending", time.Now(), time.Now(), time.Now())

	mock.ExpectQuery("INSERT INTO purchase").
		WithArgs(1, 1, "Test Store", "10000", "pending").
		WillReturnRows(rows)

	query := New(db)

	arg := CreatePurchaseParams{
		UserID:    1,
		ProductID: 1,
		StoreName: "Test Store",
		Amount:    "10000",
		Status:    "pending",
	}

	purchase, err := query.CreatePurchase(context.Background(), arg)

	require.NoError(t, err)

	require.Equal(t, purchase.UserID, arg.UserID)
}

func TestGetPurchase(t *testing.T) {
	db, mock := initTestDB(t)

	defer db.Close()

	rows := sqlmock.NewRows([]string{"purchase_id", "user_id", "product_id", "store_name", "amount", "status", "purchase_time", "created_at", "updated_at"}).
		AddRow(1, 1, 1, "Test Store", "10000", "pending", time.Now(), time.Now(), time.Now())

	mock.ExpectQuery("SELECT purchase_id, user_id, product_id, store_name, amount, status, purchase_time, created_at, updated_at FROM purchase").
		WithArgs(1).
		WillReturnRows(rows)

	query := New(db)

	purchase, err := query.GetPurchase(context.Background(), 1)

	require.NoError(t, err)

	require.Equal(t, purchase.PurchaseID, int32(1))
}

func TestGetPurchaseWithErrorNotFound(t *testing.T) {
	db, mock := initTestDB(t)

	defer db.Close()

	rows := sqlmock.NewRows([]string{"purchase_id", "user_id", "product_id", "store_name", "amount", "status", "purchase_time", "created_at", "updated_at"})

	mock.ExpectQuery("SELECT purchase_id, user_id, product_id, store_name, amount, status, purchase_time, created_at, updated_at FROM purchase").
		WithArgs(1).
		WillReturnRows(rows)

	query := New(db)

	purchase, err := query.GetPurchase(context.Background(), 1)

	require.Error(t, err)

	require.Equal(t, purchase.PurchaseID, int32(0))
}

func TestGetPurchaseWithErrorScan(t *testing.T) {
	db, mock := initTestDB(t)

	defer db.Close()

	rows := sqlmock.NewRows([]string{"purchase_id", "user_id", "product_id", "store_name", "amount", "status", "purchase_time", "created_at", "updated_at"}).
		AddRow("1", "1", "1", "Test Store", "10000", "pending", time.Now(), time.Now(), time.Now())

	mock.ExpectQuery("SELECT purchase_id, user_id, product_id, store_name, amount, status, purchase_time, created_at, updated_at FROM purchase").
		WithArgs("1").
		WillReturnRows(rows)

	query := New(db)

	purchase, err := query.GetPurchase(context.Background(), 1)

	require.Error(t, err)

	require.Equal(t, purchase.PurchaseID, int32(0))
}
