// Code generated by sqlc. DO NOT EDIT.
// source: tokens_list_by_user.sql

package db

import (
	"context"
)

const listTokensByUserID = `-- name: ListTokensByUserID :many
SELECT id, owner_id, client_id, expired_at, revoked, device_name, country_name, last_access FROM tokens WHERE
    owner_id = $1
LIMIT $2 OFFSET $3
`

type ListTokensByUserIDParams struct {
	OwnerID []byte `json:"ownerID"`
	Limit   int32  `json:"limit"`
	Offset  int32  `json:"offset"`
}

func (q *Queries) ListTokensByUserID(ctx context.Context, arg ListTokensByUserIDParams) ([]Token, error) {
	rows, err := q.query(ctx, q.listTokensByUserIDStmt, listTokensByUserID, arg.OwnerID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Token
	for rows.Next() {
		var i Token
		if err := rows.Scan(
			&i.ID,
			&i.OwnerID,
			&i.ClientID,
			&i.ExpiredAt,
			&i.Revoked,
			&i.DeviceName,
			&i.CountryName,
			&i.LastAccess,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
