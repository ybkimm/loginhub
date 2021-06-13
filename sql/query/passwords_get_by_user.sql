-- name: GetPasswordByUserID :one
SELECT * FROM passwords WHERE
    owner_id = $1
LIMIT 1;
