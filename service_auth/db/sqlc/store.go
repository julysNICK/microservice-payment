package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		Queries: New(db),
		db:      db,
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type UserTxParams struct {
	Email      string `json:"email"`
	Cpf        string `json:"cpf"`
	Password   string `json:"password"`
	CreditCard string `json:"credit_card"`
}

func (store *Store) CreateUserTx(ctx context.Context, arg UserTxParams) (int32, error) {
	var userID int32
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		userResult, err := q.CreateUser(ctx, CreateUserParams{
			Email:     arg.Email,
			Cpf:       arg.Cpf,
			Password:  arg.Password,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})

		if err != nil {
			return err
		}

		userID = userResult.ID

		_, err = q.CreateCreditCard(ctx, CreateCreditCardParams{
			UserID:    userID,
			Number:    arg.CreditCard,
			Balance:   0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})

		return err
	})
	return userID, err
}
