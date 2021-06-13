-- name: CreatePicture :exec
INSERT INTO pictures (
    id,
    desc_text,
    file_size,
    width,
    height
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
);
