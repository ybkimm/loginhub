// Code generated by sqlc. DO NOT EDIT.
// source: tokens_revoke.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const revokeToken = `-- name: RevokeToken :exec
UPDATE tokens SET
    revoked = 'true'
WHERE
    id = $1
`

func (q *Queries) RevokeToken(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.revokeTokenStmt, revokeToken, id)
	return err
}
