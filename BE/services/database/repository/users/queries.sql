-- name: CreateUser :one
INSERT INTO USERS (name, email, phone,  created_at, updated_at, deleted)
VALUES ($1, $2, $3, NOW(), NOW(), false)
ON CONFLICT (id) DO NOTHING
RETURNING id, name,  email, phone,  created_at, updated_at, deleted;


-- name: CreateCard :one
INSERT INTO Cards (card_id, user_id, device_id, active, type, assigned_at)
VALUES ($1, $2, $3, $4, $5, NOW())
ON CONFLICT (id) DO NOTHING
RETURNING id, card_id, user_id, device_id, active, type, assigned_at;

-- name: TopUpBalance :one
INSERT INTO balances (user_id, card_id, balance, ride_cost, updated_at)
VALUES ($1, $2, $3, $4,NOW())
ON CONFLICT (user_id, card_id)
DO UPDATE SET balance = EXCLUDED.balance, updated_at = NOW()
RETURNING user_id, card_id, balance, updated_at;

-- name: GetCardByCardID :one
SELECT c.id, c.card_id, c.user_id, c.active, c.type, u.name as user_name
FROM cards c
JOIN users u ON u.id = c.user_id
WHERE c.id = $1;

-- name: GetBalanceByUserID :one
SELECT user_id, balance, ride_cost, updated_at
FROM balances
WHERE user_id = $1;

-- name: DeductBalance :exec
UPDATE balances SET balance = balance - $1, updated_at = NOW() WHERE user_id = $2 AND balance >= $1;

-- name: InsertCharge :exec
INSERT INTO charges (user_id, amount, type, description, created_at) VALUES ($1, $2, 'ride', $3, NOW());


-- name: CountUsers :one
SELECT COUNT(*) FROM users;

-- name: CountCards :one
SELECT COUNT(*) FROM cards;

-- name: TotalBalance :one
SELECT COALESCE(SUM(balance), 0) FROM balances;

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

-- name: CardsList :many
SELECT
  c.id,
  c.card_id,
  c.user_id,
  c.device_id,
  c.type,
  c.active,
  c.assigned_at,
  COUNT(*) OVER() AS total
FROM cards c
JOIN users u ON u.id = c.user_id
WHERE u.deleted = false
ORDER BY c.id
LIMIT $1 OFFSET $2;

-- name: ChargesList :many
SELECT
  c.*,
  COUNT(*) OVER() AS total
FROM charges c
ORDER BY c.id
LIMIT $1 OFFSET $2;

-- name: CostOfRide :exec
UPDATE balances
SET
    ride_cost = $1,
    updated_at = NOW();


-- name: BalaneList :many
SELECT
  b.*,
  COUNT(*) OVER() AS total
FROM balances b
ORDER BY b.id
LIMIT $1 OFFSET $2;


-- name: DeviceList :many
SELECT
  b.*,
  COUNT(*) OVER() AS total
FROM devices b
ORDER BY b.id
LIMIT $1 OFFSET $2;

-- name: CreateDevice :one
INSERT INTO devices (device_id, location, active, installed_at)
VALUES ($1, $2, true, NOW())
ON CONFLICT (id) DO NOTHING
RETURNING *;

-- name: CreateCardActivation :one
INSERT INTO card_activations (card_id, activation_start, activation_end, created_at, updated_at)
VALUES ($1, $2, $3, NOW(), NOW())
RETURNING id, card_id, activation_start, activation_end, created_at, updated_at;

-- name: GetUserByID :one
SELECT id, name, email, phone, deleted FROM users WHERE id = $1;

-- name: UpdateUser :exec
UPDATE users SET name = $2, email = $3, phone = $4, updated_at = NOW() WHERE id = $1;

-- name: SoftDeleteUser :exec
UPDATE users SET deleted = TRUE WHERE id = $1;

-- name: GetCardByID :one
SELECT id, card_id, user_id, device_id, active, type, assigned_at FROM cards WHERE id = $1;

-- name: UpdateCard :exec
UPDATE cards SET card_id = $2, user_id = $3, device_id = $4, type = $5, active = $6 WHERE id = $1;

-- name: SoftDeleteCard :exec
UPDATE cards SET deleted = TRUE WHERE id = $1;

-- name: GetDeviceByID :one
SELECT id, device_id, location, installed_at, active FROM devices WHERE id = $1;

-- name: UpdateDevice :exec
UPDATE devices SET device_id = $2, location = $3, active = $4 WHERE id = $1;

-- name: SoftDeleteDevice :exec
UPDATE devices SET active = FALSE WHERE id = $1;

-- name: GetAuthorizedAccessByDeviceUniqueID :many
SELECT
    c.card_id,
    u.id AS user_id,
    u.name AS user_name,
    c.type,
    c.active,
    b.balance,
    b.ride_cost,
    a.activation_start,
    a.activation_end
FROM cards c
JOIN users u ON u.id = c.user_id
LEFT JOIN balances b ON b.card_id = c.id
LEFT JOIN LATERAL (
    SELECT activation_start, activation_end
    FROM card_activations
    WHERE card_id = c.id
    ORDER BY activation_end DESC
    LIMIT 1
) a ON TRUE
JOIN devices d ON d.id = c.device_id
WHERE d.device_id = $1 AND c.deleted = FALSE;


-- name: GetSumBalanceByDeviceID :one
SELECT COALESCE(SUM(b.balance), 0) AS total_balance
FROM balances b
JOIN cards c ON b.card_id = c.id
JOIN devices d ON c.device_id = d.id
WHERE d.device_id = $1;


-- AUTH USERS QUERIES

-- name: CreateAuthUser :one
INSERT INTO auth_users (username, password_hash, role)
VALUES ($1, $2, $3)
RETURNING id, username, password_hash, role, created_at, updated_at;

-- name: GetAuthUserByUsername :one
SELECT id, username, password_hash, role, device_id, created_at, updated_at
FROM auth_users
WHERE username = $1;

-- name: GetAuthUserByID :one
SELECT id, username, password_hash, role, created_at, updated_at
FROM auth_users
WHERE id = $1;

-- name: ListAuthUsers :many
SELECT id, username, password_hash, role, created_at, updated_at
FROM auth_users
ORDER BY id; 