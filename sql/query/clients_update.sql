-- name: UpdateClient :exec
UPDATE clients SET
    client_name = $2,
    client_category = $3,
    client_desc = $4,
    banner_picture_id = $5,
    icon_id = $6,
    redirect_uris = $7,
    terms_url = $8,
    privacy_policy_url = $9,
    flag = $10
WHERE
    id = $1;
