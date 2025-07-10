-- name: CreateUser :one
INSERT INTO USERS (name, email, phone,  created_at, updated_at, deleted)
VALUES ($1, $2, $3, NOW(), NOW(), false)
ON CONFLICT (id) DO NOTHING
RETURNING id, name,  email, phone,  created_at, updated_at, deleted;

-- name: UsersList :many
SELECT
  u.id,
  u.name,
  u.email,
  u.phone,
  COALESCE(card_counts.card_count, 0) AS card_count,
 COALESCE(balance_sums.total_balance, 0)::float8 AS total_balance,
  COUNT(*) OVER() AS total
FROM users u
LEFT JOIN (
    SELECT user_id, COUNT(*) AS card_count
    FROM cards
    GROUP BY user_id
) AS card_counts ON card_counts.user_id = u.id
LEFT JOIN (
    SELECT user_id, SUM(balance) AS total_balance
    FROM balances
    GROUP BY user_id
) AS balance_sums ON balance_sums.user_id = u.id
WHERE deleted = false
ORDER BY u.id
LIMIT $1 OFFSET $2;

-- name: GetUserByID :one
SELECT id, name, email, phone, deleted FROM users WHERE id = $1;

-- name: UpdateUser :exec
UPDATE users SET name = $2, email = $3, phone = $4, updated_at = NOW() WHERE id = $1;

-- name: SoftDeleteUser :exec
UPDATE users SET deleted = TRUE WHERE id = $1;

-- name: CountUsers :one
SELECT COUNT(*) FROM users;