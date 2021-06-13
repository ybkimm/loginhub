-- name: RevokeToken :exec
UPDATE FROM tokens SET
    revoked = 'true'
WHERE
    id = $1;
