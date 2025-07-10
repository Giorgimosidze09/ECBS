-- name: TopUpBalance :one
INSERT INTO balances (user_id, card_id, balance, ride_cost, updated_at)
VALUES ($1, $2, $3, $4,NOW())
ON CONFLICT (user_id, card_id)
DO UPDATE SET balance = EXCLUDED.balance, updated_at = NOW()
RETURNING user_id, card_id, balance, updated_at;

-- name: GetBalanceByUserID :one
SELECT user_id, balance, ride_cost, updated_at
FROM balances
WHERE user_id = $1;

-- name: DeductBalance :exec
UPDATE balances SET balance = balance - $1, updated_at = NOW() WHERE user_id = $2 AND balance >= $1;

-- name: BalaneList :many
SELECT
  b.*,
  COUNT(*) OVER() AS total
FROM balances b
ORDER BY b.id
LIMIT $1 OFFSET $2;

-- name: GetSumBalanceByDeviceID :one
SELECT COALESCE(SUM(b.balance), 0) AS total_balance
FROM balances b
JOIN cards c ON b.card_id = c.id
JOIN devices d ON c.device_id = d.id
WHERE d.device_id = $1;

-- name: AddBalanceToCard :one
UPDATE balances
SET balance = balance + $2, updated_at = NOW()
WHERE card_id = $1
RETURNING user_id, card_id, balance, updated_at;

-- name: TotalBalance :one
SELECT COALESCE(SUM(balance), 0) FROM balances;

-- name: CostOfRide :exec
UPDATE balances
SET
    ride_cost = $1,
    updated_at = NOW(); 