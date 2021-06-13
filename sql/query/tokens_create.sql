-- name: CreateToken :exec
INSERT INTO tokens (
    id,
    owner_id,
    client_id,
    expired_at,
    device_name,
    country_name
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
);
