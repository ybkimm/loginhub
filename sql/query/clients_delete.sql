-- name: DeleteClient :exec
DELETE FROM clients
    id = $1;
