// Code generated by sqlc. DO NOT EDIT.
// source: tokens_create.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createToken = `-- name: CreateToken :exec
INSERT INTO tokens (
    id,
    owner_id,
    client_id,
    expired_at,
    device_name,
    country_name
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
`

type CreateTokenParams struct {
	ID          uuid.UUID `json:"id"`
	OwnerID     uuid.UUID `json:"ownerID"`
	ClientID    uuid.UUID `json:"clientID"`
	ExpiredAt   time.Time `json:"expiredAt"`
	DeviceName  string    `json:"deviceName"`
	CountryName string    `json:"countryName"`
}

func (q *Queries) CreateToken(ctx context.Context, arg CreateTokenParams) error {
	_, err := q.exec(ctx, q.createTokenStmt, createToken,
		arg.ID,
		arg.OwnerID,
		arg.ClientID,
		arg.ExpiredAt,
		arg.DeviceName,
		arg.CountryName,
	)
	return err
}
