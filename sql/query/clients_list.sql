-- name: ListClientAll :many
SELECT * FROM clients LIMIT $2 OFFSET $3;
