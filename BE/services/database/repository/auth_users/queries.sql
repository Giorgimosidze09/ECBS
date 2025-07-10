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