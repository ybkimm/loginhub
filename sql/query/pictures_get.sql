-- name: GetPicture :one
SELECT * FROM pictures WHERE
    id = $1
LIMIT 1;
