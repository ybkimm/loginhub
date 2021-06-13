-- name: RevokePassword :exec
UPDATE passwords SET
    revoked = 'true'
WHERE
    owner_id = $1 AND
    revoked = 'false';
