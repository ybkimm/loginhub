-- name: CreatePassword :exec
INSERT INTO passwords (
    id,
    owner_id,
    pass_hash
) VALUES (
    $1,
    $2,
    $3
);
