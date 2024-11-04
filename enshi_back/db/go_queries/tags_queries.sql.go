// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: tags_queries.sql

package db_repo

import (
	"context"
)

const createTag = `-- name: CreateTag :one
INSERT INTO public.tags
(tag_id, tag_name)
VALUES($1, $2)
RETURNING tag_id, tag_name
`

type CreateTagParams struct {
	TagID   int32  `json:"tag_id"`
	TagName string `json:"tag_name"`
}

func (q *Queries) CreateTag(ctx context.Context, arg CreateTagParams) (Tag, error) {
	row := q.db.QueryRow(ctx, createTag, arg.TagID, arg.TagName)
	var i Tag
	err := row.Scan(&i.TagID, &i.TagName)
	return i, err
}

const deleteTag = `-- name: DeleteTag :exec
DELETE FROM public.tags
WHERE tag_id = $1
`

func (q *Queries) DeleteTag(ctx context.Context, tagID int32) error {
	_, err := q.db.Exec(ctx, deleteTag, tagID)
	return err
}

const getAllTags = `-- name: GetAllTags :many
SELECT tag_id, tag_name
FROM public.tags
ORDER BY tag_name ASC
`

func (q *Queries) GetAllTags(ctx context.Context) ([]Tag, error) {
	rows, err := q.db.Query(ctx, getAllTags)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tag
	for rows.Next() {
		var i Tag
		if err := rows.Scan(&i.TagID, &i.TagName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTagByTagId = `-- name: GetTagByTagId :one
SELECT tag_id, tag_name
FROM public.tags tags
where tags.tag_id = $1
`

func (q *Queries) GetTagByTagId(ctx context.Context, tagID int32) (Tag, error) {
	row := q.db.QueryRow(ctx, getTagByTagId, tagID)
	var i Tag
	err := row.Scan(&i.TagID, &i.TagName)
	return i, err
}
