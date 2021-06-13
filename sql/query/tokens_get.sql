-- name: GetToken :one
SELECT * FROM tokens WHERE
    id = $1
LIMIT 1;
