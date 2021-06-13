-- name: ListClientByUserID :many
SELECT * FROM clients WHERE
    owner_id = $1
LIMIT $2 OFFSET $3;
