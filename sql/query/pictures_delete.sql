-- name: DeletePicture :exec
DELETE FROM pictures WHERE
    id = $1;
