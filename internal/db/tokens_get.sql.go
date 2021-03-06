// Code generated by sqlc. DO NOT EDIT.
// source: tokens_get.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const getToken = `-- name: GetToken :one
SELECT id, owner_id, client_id, expired_at, revoked, device_name, country_name, last_access FROM tokens WHERE
    id = $1
LIMIT 1
`

func (q *Queries) GetToken(ctx context.Context, id uuid.UUID) (Token, error) {
	row := q.queryRow(ctx, q.getTokenStmt, getToken, id)
	var i Token
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.ClientID,
		&i.ExpiredAt,
		&i.Revoked,
		&i.DeviceName,
		&i.CountryName,
		&i.LastAccess,
	)
	return i, err
}
