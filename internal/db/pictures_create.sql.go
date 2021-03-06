// Code generated by sqlc. DO NOT EDIT.
// source: pictures_create.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createPicture = `-- name: CreatePicture :exec
INSERT INTO pictures (
    id,
    desc_text,
    file_size,
    width,
    height
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
`

type CreatePictureParams struct {
	ID       uuid.UUID `json:"id"`
	DescText string    `json:"descText"`
	FileSize int32     `json:"fileSize"`
	Width    int32     `json:"width"`
	Height   int32     `json:"height"`
}

func (q *Queries) CreatePicture(ctx context.Context, arg CreatePictureParams) error {
	_, err := q.exec(ctx, q.createPictureStmt, createPicture,
		arg.ID,
		arg.DescText,
		arg.FileSize,
		arg.Width,
		arg.Height,
	)
	return err
}
