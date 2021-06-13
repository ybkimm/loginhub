-- name: CreateUser :exec
INSERT INTO users (
    id,
    email,
    picture_id,
    given_name,
    family_name,
    gender,
    birthdate,
    flag
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
);
