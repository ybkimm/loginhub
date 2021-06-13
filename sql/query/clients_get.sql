-- name: GetClient :one
SELECT * FROM clients WHERE
    id = $1
LIMIT 1;
