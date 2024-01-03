-- name: CreateUser :one
INSERT INTO users (email, password, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id, email, password, created_at, updated_at;

-- name: GetUserByUsername :one
SELECT id, password, email, created_at, updated_at FROM users WHERE email = $1 LIMIT 1;
