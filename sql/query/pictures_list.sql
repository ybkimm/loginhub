-- name: ListPictures :many
SELECT * FROM pictures
LIMIT $1 OFFSET $2;
