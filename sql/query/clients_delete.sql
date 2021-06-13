-- name: DeleteClient :exec
DELETE FROM clients WHERE
    id = $1;
