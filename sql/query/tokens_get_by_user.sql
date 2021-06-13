-- name: GetTokenByUserID :one
SELECT * FROM tokens WHERE
    owner_id = $1
LIMIT 1;
