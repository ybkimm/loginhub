-- name: ListUsers :many
SELECT * FROM users LIMIT $1 OFFSET $2;
