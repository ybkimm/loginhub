-- name: ListTokensByUserID :many
SELECT * FROM tokens WHERE
    owner_id = $1
LIMIT $2 OFFSET $3;
