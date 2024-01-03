-- name: CreateCreditCard :one
INSERT INTO credit_card (user_id, number, balance, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, user_id, number, balance, created_at, updated_at;

-- name: GetCreditCardByUserID :one
SELECT id, user_id, number, balance, created_at, updated_at FROM credit_card WHERE user_id = $1 LIMIT 1;