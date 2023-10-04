// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: tags.sql

package queries

import (
	"context"
)

const createTag = `-- name: CreateTag :one
INSERT INTO tags (
  user_id, 
  name,
  sign,
  kind
) VALUES (
  $1, 
  $2, 
  $3, 
  $4
) RETURNING id, user_id, name, sign, kind, deleted_at, created_at, updated_at
`

type CreateTagParams struct {
	UserID int32  `json:"user_id"`
	Name   string `json:"name"`
	Sign   string `json:"sign"`
	Kind   string `json:"kind"`
}

func (q *Queries) CreateTag(ctx context.Context, arg CreateTagParams) (Tag, error) {
	row := q.db.QueryRowContext(ctx, createTag,
		arg.UserID,
		arg.Name,
		arg.Sign,
		arg.Kind,
	)
	var i Tag
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Sign,
		&i.Kind,
		&i.DeletedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteTag = `-- name: DeleteTag :exec
UPDATE tags
SET deleted_at = now()
WHERE id = $1
`

func (q *Queries) DeleteTag(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteTag, id)
	return err
}

const findTag = `-- name: FindTag :one
SELECT id, user_id, name, sign, kind, deleted_at, created_at, updated_at FROM tags
WHERE id = $1 AND user_id = $2 AND deleted_at IS NULL
`

type FindTagParams struct {
	ID     int32 `json:"id"`
	UserID int32 `json:"user_id"`
}

func (q *Queries) FindTag(ctx context.Context, arg FindTagParams) (Tag, error) {
	row := q.db.QueryRowContext(ctx, findTag, arg.ID, arg.UserID)
	var i Tag
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Sign,
		&i.Kind,
		&i.DeletedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listTags = `-- name: ListTags :many
SELECT id, user_id, name, sign, kind, deleted_at, created_at, updated_at FROM tags
WHERE kind = $3 AND user_id = $4 AND deleted_at IS NULL
ORDER BY created_at DESC
OFFSET $1
LIMIT $2
`

type ListTagsParams struct {
	Offset int32  `json:"offset"`
	Limit  int32  `json:"limit"`
	Kind   string `json:"kind"`
	UserID int32  `json:"user_id"`
}

func (q *Queries) ListTags(ctx context.Context, arg ListTagsParams) ([]Tag, error) {
	rows, err := q.db.QueryContext(ctx, listTags,
		arg.Offset,
		arg.Limit,
		arg.Kind,
		arg.UserID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tag
	for rows.Next() {
		var i Tag
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
			&i.Sign,
			&i.Kind,
			&i.DeletedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateTag = `-- name: UpdateTag :one
UPDATE tags
SET
  user_id = $1,
  name = CASE WHEN $2::varchar = '' THEN name ELSE $2 END,
  sign = CASE WHEN $3::varchar = '' THEN sign ELSE $3 END,
  kind = CASE WHEN $4::varchar = '' THEN kind ELSE $4 END
WHERE id = $5
RETURNING id, user_id, name, sign, kind, deleted_at, created_at, updated_at
`

type UpdateTagParams struct {
	UserID int32  `json:"user_id"`
	Name   string `json:"name"`
	Sign   string `json:"sign"`
	Kind   string `json:"kind"`
	ID     int32  `json:"id"`
}

func (q *Queries) UpdateTag(ctx context.Context, arg UpdateTagParams) (Tag, error) {
	row := q.db.QueryRowContext(ctx, updateTag,
		arg.UserID,
		arg.Name,
		arg.Sign,
		arg.Kind,
		arg.ID,
	)
	var i Tag
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Sign,
		&i.Kind,
		&i.DeletedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
