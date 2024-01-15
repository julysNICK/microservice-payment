-- name: CreateCreditCard :one
INSERT INTO credit_card (user_id, number, balance, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, user_id, number, balance, created_at, updated_at;

-- name: GetCreditCardByUserID :many
SELECT id, user_id, number, balance, created_at, updated_at FROM credit_card WHERE user_id = $1 LIMIT 1;

-- name: GetCreditCardByID :many
SELECT id, user_id, number, balance, created_at, updated_at FROM credit_card WHERE id = $1 LIMIT 1;

-- name: UpdateCreditCardPartial :exec
UPDATE credit_card SET number = COALESCE(sqlc.narg(number), number), balance = COALESCE(sqlc.narg(balance), balance), updated_at = COALESCE(sqlc.narg(updated_at), updated_at) WHERE id = $1;