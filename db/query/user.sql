-- name: CreateUser :one
INSERT INTO users (
        first_name,
        last_name,
        bio,
        username,
        password,
        email
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;
-- name: ListUsers :many
SELECT *
FROM users
ORDER BY id
LIMIT $1 OFFSET $2;
-- name: UpdateUserBio :one
UPDATE users
SET bio = $1
WHERE id = $1
RETURNING *;
-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;