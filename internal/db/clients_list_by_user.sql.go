// Code generated by sqlc. DO NOT EDIT.
// source: clients_list_by_user.sql

package db

import (
	"context"
)

const listClientByUserID = `-- name: ListClientByUserID :many
SELECT id, client_name, secret_hash, owner_id, client_category, client_desc, banner_picture_id, icon_id, redirect_uris, terms_url, privacy_policy_url, flag FROM clients WHERE
    owner_id = $1
LIMIT $2 OFFSET $3
`

type ListClientByUserIDParams struct {
	OwnerID []byte `json:"ownerID"`
	Limit   int32  `json:"limit"`
	Offset  int32  `json:"offset"`
}

func (q *Queries) ListClientByUserID(ctx context.Context, arg ListClientByUserIDParams) ([]Client, error) {
	rows, err := q.query(ctx, q.listClientByUserIDStmt, listClientByUserID, arg.OwnerID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Client
	for rows.Next() {
		var i Client
		if err := rows.Scan(
			&i.ID,
			&i.ClientName,
			&i.SecretHash,
			&i.OwnerID,
			&i.ClientCategory,
			&i.ClientDesc,
			&i.BannerPictureID,
			&i.IconID,
			&i.RedirectUris,
			&i.TermsUrl,
			&i.PrivacyPolicyUrl,
			&i.Flag,
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