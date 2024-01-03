// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"time"
)

type CreditCard struct {
	ID        int32
	UserID    int32
	Number    string
	Balance   int32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID        int32
	Email     string
	Cpf       string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}