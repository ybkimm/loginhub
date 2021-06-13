-- name: CreateClient :exec
INSERT INTO clients (
    id,
    client_name,
    secret_hash,
    owner_id,
    client_category,
    client_desc,
    banner_picture_id,
    icon_id,
    redirect_uris,
    terms_url,
    privacy_policy_url,
    flag
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11,
    $12
);
