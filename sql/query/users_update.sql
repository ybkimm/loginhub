-- name: UpdateUser :exec
UPDATE users SET
    email = $2,
    picture_id = $3,
    given_name = $4,
    family_name = $5,
    gender = $6,
    birthdate = $7,
    flag = $8
WHERE
    id = $1;
