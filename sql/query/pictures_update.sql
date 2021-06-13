-- name: UpdatePicture :exec
UPDATE pictures SET
    desc_text = $2,
    file_size = $3,
    width = $4,
    height = $5
WHERE
    id = $1;
