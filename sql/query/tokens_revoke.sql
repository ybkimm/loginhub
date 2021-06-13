-- name: RevokeToken :exec
UPDATE tokens SET
    revoked = 'true'
WHERE
    id = $1;
