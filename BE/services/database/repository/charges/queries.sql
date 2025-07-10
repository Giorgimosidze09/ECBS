-- name: InsertCharge :exec
INSERT INTO charges (user_id, amount, type, description, created_at) VALUES ($1, $2, 'ride', $3, NOW());

-- name: ChargesList :many
SELECT
  c.*,
  COUNT(*) OVER() AS total
FROM charges c
ORDER BY c.id
LIMIT $1 OFFSET $2; 