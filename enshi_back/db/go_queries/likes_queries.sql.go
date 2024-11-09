// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: likes_queries.sql

package db_repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createLike = `-- name: CreateLike :one
INSERT INTO public.likes
(like_id, user_id, comment_id, created_at)
VALUES($1, $2, $3, CURRENT_TIMESTAMP)
RETURNING like_id, user_id, comment_id, created_at
`

type CreateLikeParams struct {
	LikeID    int64       `json:"like_id"`
	UserID    pgtype.Int8 `json:"user_id"`
	CommentID pgtype.Int8 `json:"comment_id"`
}

func (q *Queries) CreateLike(ctx context.Context, arg CreateLikeParams) (Like, error) {
	row := q.db.QueryRow(ctx, createLike, arg.LikeID, arg.UserID, arg.CommentID)
	var i Like
	err := row.Scan(
		&i.LikeID,
		&i.UserID,
		&i.CommentID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteLikeByUserCommentId = `-- name: DeleteLikeByUserCommentId :exec
DELETE FROM public.likes
WHERE user_id = $1 and comment_id = $2
`

type DeleteLikeByUserCommentIdParams struct {
	UserID    pgtype.Int8 `json:"user_id"`
	CommentID pgtype.Int8 `json:"comment_id"`
}

func (q *Queries) DeleteLikeByUserCommentId(ctx context.Context, arg DeleteLikeByUserCommentIdParams) error {
	_, err := q.db.Exec(ctx, deleteLikeByUserCommentId, arg.UserID, arg.CommentID)
	return err
}

const getLikesForComment = `-- name: GetLikesForComment :one
SELECT count(*)
FROM public.likes
WHERE comment_id = $1
`

func (q *Queries) GetLikesForComment(ctx context.Context, commentID pgtype.Int8) (int64, error) {
	row := q.db.QueryRow(ctx, getLikesForComment, commentID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const isUserLikedComment = `-- name: IsUserLikedComment :one
SELECT count(*)
FROM public.likes
WHERE user_id = $1 and comment_id = $2
`

type IsUserLikedCommentParams struct {
	UserID    pgtype.Int8 `json:"user_id"`
	CommentID pgtype.Int8 `json:"comment_id"`
}

func (q *Queries) IsUserLikedComment(ctx context.Context, arg IsUserLikedCommentParams) (int64, error) {
	row := q.db.QueryRow(ctx, isUserLikedComment, arg.UserID, arg.CommentID)
	var count int64
	err := row.Scan(&count)
	return count, err
}
