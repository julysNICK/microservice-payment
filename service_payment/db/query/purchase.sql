-- name: CreatePurchase :one
INSERT INTO purchase (
  user_id,
  product_id,
  store_name,
  amount,
  status
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
) RETURNING purchase_id, user_id, product_id, store_name, amount, status, purchase_time, created_at, updated_at;

-- name: GetPurchase :one
SELECT purchase_id, user_id, product_id, store_name, amount, status, purchase_time, created_at, updated_at FROM purchase WHERE purchase_id = $1;