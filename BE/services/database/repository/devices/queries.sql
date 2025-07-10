-- name: CreateDevice :one
INSERT INTO devices (device_id, location, active, installed_at)
VALUES ($1, $2, true, NOW())
ON CONFLICT (id) DO NOTHING
RETURNING *;

-- name: GetDeviceByID :one
SELECT id, device_id, location, installed_at, active FROM devices WHERE id = $1;

-- name: UpdateDevice :exec
UPDATE devices SET device_id = $2, location = $3, active = $4 WHERE id = $1;

-- name: SoftDeleteDevice :exec
UPDATE devices SET active = FALSE WHERE id = $1;

-- name: DeviceList :many
SELECT
  b.*,
  COUNT(*) OVER() AS total
FROM devices b
ORDER BY b.id
LIMIT $1 OFFSET $2;

-- name: GetAuthorizedAccessByDeviceUniqueID :many
SELECT
    c.card_id,
    u.id AS user_id,
    u.name AS user_name,
    u.pin_code, 
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