-- name: CreateCard :one
INSERT INTO Cards (card_id, user_id, device_id, active, type, assigned_at)
VALUES ($1, $2, $3, $4, $5, NOW())
ON CONFLICT (id) DO NOTHING
RETURNING id, card_id, user_id, device_id, active, type, assigned_at;

-- name: GetCardByCardID :one
SELECT c.id, c.card_id, c.user_id, c.active, c.type, u.name as user_name
FROM cards c
JOIN users u ON u.id = c.user_id
WHERE c.id = $1;

-- name: GetCardByID :one
SELECT id, card_id, user_id, device_id, active, type, assigned_at FROM cards WHERE id = $1;

-- name: UpdateCard :exec
UPDATE cards SET card_id = $2, user_id = $3, device_id = $4, type = $5, active = $6 WHERE id = $1;

-- name: SoftDeleteCard :exec
UPDATE cards SET deleted = TRUE WHERE id = $1;

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

-- name: CreateCardActivation :one
INSERT INTO card_activations (card_id, activation_start, activation_end, created_at, updated_at)
VALUES ($1, $2, $3, NOW(), NOW())
RETURNING id, card_id, activation_start, activation_end, created_at, updated_at;

-- name: GetCardByItsCardID :one
SELECT c.id, c.card_id, c.user_id, c.active, c.type, u.name as user_name
FROM cards c
JOIN users u ON u.id = c.user_id
WHERE c.card_id = $1;

-- name: CountCards :one
SELECT COUNT(*) FROM cards; 