// Code generated by sqlc. DO NOT EDIT.
// source: clients_create.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createClient = `-- name: CreateClient :exec
INSERT INTO clients (
    id,
    client_name,
    secret_hash,
    owner_id,
    client_category,
    client_desc,
    banner_picture_id,
    icon_id,
    redirect_uris,
    terms_url,
    privacy_policy_url,
    flag
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11,
    $12
)
`

type CreateClientParams struct {
	ID               uuid.UUID      `json:"id"`
	ClientName       string         `json:"clientName"`
	SecretHash       string         `json:"secretHash"`
	OwnerID          uuid.UUID      `json:"ownerID"`
	ClientCategory   string         `json:"clientCategory"`
	ClientDesc       string         `json:"clientDesc"`
	BannerPictureID  uuid.UUID      `json:"bannerPictureID"`
	IconID           uuid.UUID      `json:"iconID"`
	RedirectUris     sql.NullString `json:"redirectUris"`
	TermsUrl         sql.NullString `json:"termsUrl"`
	PrivacyPolicyUrl sql.NullString `json:"privacyPolicyUrl"`
	Flag             sql.NullInt32  `json:"flag"`
}

func (q *Queries) CreateClient(ctx context.Context, arg CreateClientParams) error {
	_, err := q.exec(ctx, q.createClientStmt, createClient,
		arg.ID,
		arg.ClientName,
		arg.SecretHash,
		arg.OwnerID,
		arg.ClientCategory,
		arg.ClientDesc,
		arg.BannerPictureID,
		arg.IconID,
		arg.RedirectUris,
		arg.TermsUrl,
		arg.PrivacyPolicyUrl,
		arg.Flag,
	)
	return err
}
