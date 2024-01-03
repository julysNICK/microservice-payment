package data

// import (
// 	"database/sql"
// 	"service_auth/data"

// 	"time"
// )

// const dbTimeout = time.Second * 3

// var db *sql.DB

// func New(dbPool *sql.DB) Models {
// 	db = dbPool

// 	return Models{
// 		User: data.UserModel{},
// 	}
// }

// type Models struct {
// 	User data.UserModel
// }

// func (u *data.UserModel) GetUserByEmail(email string) (*data.UserModel, error) {

// 	query := `SELECT id, email, password FROM users WHERE email = $1`

// 	row := db.QueryRow(query, email)

// 	err := row.Scan(&u.ID, &u.Email, &u.Password)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return u, nil
// }

// func (u *data.UserModel) GetUserByID(id int) (*data.UserModel, error) {

// 	query := `SELECT id, email, password FROM users WHERE id = $1`

// 	row := db.QueryRow(query, id)

// 	err := row.Scan(&u.ID, &u.Email, &u.Password)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return u, nil
// }
