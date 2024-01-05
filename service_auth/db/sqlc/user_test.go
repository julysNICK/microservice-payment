package db

import (
	"context"
	"fmt"
	"service_auth/util/mocks"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Email:     mocks.RandomEmail(),
		Cpf:       mocks.GenerateRandomCPF(),
		Password:  mocks.GenerateRandomPassword(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	fmt.Println(arg.Cpf)

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)

	require.NotEmpty(t, user)

	require.Equal(t, arg.Email, user.Email)

	require.Equal(t, arg.Cpf, user.Cpf)

	require.Equal(t, arg.Password, user.Password)

	return user

}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetUserByUsername(t *testing.T) {
	user := CreateRandomUser(t)

	user2, err := testQueries.GetUserByUsername(context.Background(), user.Email)

	require.NoError(t, err)

	require.NotEmpty(t, user2)

	require.Equal(t, user.Email, user2.Email)

	require.Equal(t, user.Cpf, user2.Cpf)

	require.Equal(t, user.Password, user2.Password)
}

// func TestGetUserByUsernameWithMock(t *testing.T) {
// 	db, mock, err := sqlmock.New()

// 	require.NoError(t, err)

// 	defer db.Close()

// }

func TestInsertUserWithMock(t *testing.T) {
	db, mock, err := sqlmock.New()

	require.NoError(t, err)

	defer db.Close()

	arg := CreateUserParams{
		Email:     mocks.RandomEmail(),
		Cpf:       mocks.GenerateRandomCPF(),
		Password:  mocks.GenerateRandomPassword(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "email", "cpf", "password", "created_at", "updated_at"}).
		AddRow(1, arg.Email, arg.Cpf, arg.Password, arg.CreatedAt, arg.UpdatedAt)

	mock.ExpectQuery("INSERT INTO users").
		WithArgs(arg.Email, arg.Cpf, arg.Password, arg.CreatedAt, arg.UpdatedAt).
		WillReturnRows(rows)

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)

	require.NotEmpty(t, user)

	require.Equal(t, arg.Email, user.Email)

	require.Equal(t, arg.Cpf, user.Cpf)

	require.Equal(t, arg.Password, user.Password)
}

func TestGetUserByUsernameWithMockSuccess(t *testing.T) {
	users := []User{
		{
			ID:        1,
			Email:     mocks.RandomEmail(),
			Cpf:       mocks.GenerateRandomCPF(),
			Password:  mocks.GenerateRandomPassword(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Email:     mocks.RandomEmail(),
			Cpf:       mocks.GenerateRandomCPF(),
			Password:  mocks.GenerateRandomPassword(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        3,
			Email:     mocks.RandomEmail(),
			Cpf:       mocks.GenerateRandomCPF(),
			Password:  mocks.GenerateRandomPassword(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	db, mock, err := sqlmock.New()

	require.NoError(t, err)

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "email", "cpf", "password", "created_at", "updated_at"}).AddRow(users[0].ID, users[0].Email, users[0].Cpf, users[0].Password, users[0].CreatedAt, users[0].UpdatedAt)

	mock.ExpectQuery("SELECT (.+) FROM users WHERE email = (.+)").WithArgs(users[0].Email).WillReturnRows(rows)

	db.QueryRow("SELECT (.+) FROM users WHERE email = (.+)", users[0].Email).Scan(&users[0].ID, &users[0].Email, &users[0].Cpf, &users[0].Password, &users[0].CreatedAt, &users[0].UpdatedAt)

	require.NoError(t, err)

	require.NotEmpty(t, users[0])

	require.Equal(t, users[0].Email, users[0].Email)

	require.Equal(t, users[0].Cpf, users[0].Cpf)

	require.Equal(t, users[0].Password, users[0].Password)

}

func TestGetUserByUsernameWithMockErrorNotFound(t *testing.T) {
	users := []User{
		{
			ID:        1,
			Email:     mocks.RandomEmail(),
			Cpf:       mocks.GenerateRandomCPF(),
			Password:  mocks.GenerateRandomPassword(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Email:     mocks.RandomEmail(),
			Cpf:       mocks.GenerateRandomCPF(),
			Password:  mocks.GenerateRandomPassword(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        3,
			Email:     mocks.RandomEmail(),
			Cpf:       mocks.GenerateRandomCPF(),
			Password:  mocks.GenerateRandomPassword(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	db, mock, err := sqlmock.New()

	require.NoError(t, err)

	defer db.Close()

	mockDb := New(db)

	mock.ExpectQuery("SELECT (.+) FROM users WHERE email = (.+)").WithArgs(users[0].Email).WillReturnError(fmt.Errorf("error"))

	_, err = mockDb.GetUserByUsername(context.Background(), users[0].Email)

	require.Error(t, err)

	require.EqualError(t, err, "error")

}

func TestGetUserByUsernameWithMockRowsErr(t *testing.T) {
	db, mock, err := sqlmock.New()

	require.NoError(t, err)

	defer db.Close()

	mockDb := New(db)

	rows := sqlmock.NewRows([]string{"id", "email", "cpf", "password", "created_at", "updated_at"}).AddRow(1, "email", "cpf", "password", time.Now(), time.Now()).CloseError(fmt.Errorf("error"))

	mock.ExpectQuery("SELECT (.+) FROM users WHERE email = (.+)").WithArgs("email").WillReturnRows(rows)

	_, err = mockDb.GetUserByUsername(context.Background(), "email")

	require.Error(t, err)

	require.EqualError(t, err, "error")
}
