-- name: RevokeAllTokensByUserID :exec
UPDATE FROM passwords SET
    revoked = 'true'
WHERE
    owner_id = $1 AND
    revoked = 'false';
