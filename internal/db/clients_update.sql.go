// Code generated by sqlc. DO NOT EDIT.
// source: clients_update.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const updateClient = `-- name: UpdateClient :exec
UPDATE clients SET
    client_name = $2,
    client_category = $3,
    client_desc = $4,
    banner_picture_id = $5,
    icon_id = $6,
    redirect_uris = $7,
    terms_url = $8,
    privacy_policy_url = $9,
    flag = $10
WHERE
    id = $1
`

type UpdateClientParams struct {
	ID               uuid.UUID      `json:"id"`
	ClientName       string         `json:"clientName"`
	ClientCategory   string         `json:"clientCategory"`
	ClientDesc       string         `json:"clientDesc"`
	BannerPictureID  uuid.UUID      `json:"bannerPictureID"`
	IconID           uuid.UUID      `json:"iconID"`
	RedirectUris     sql.NullString `json:"redirectUris"`
	TermsUrl         sql.NullString `json:"termsUrl"`
	PrivacyPolicyUrl sql.NullString `json:"privacyPolicyUrl"`
	Flag             sql.NullInt32  `json:"flag"`
}

func (q *Queries) UpdateClient(ctx context.Context, arg UpdateClientParams) error {
	_, err := q.exec(ctx, q.updateClientStmt, updateClient,
		arg.ID,
		arg.ClientName,
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
