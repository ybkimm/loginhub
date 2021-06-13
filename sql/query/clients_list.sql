-- name: ListClientAll :many
SELECT * FROM clients LIMIT $1 OFFSET $2;
