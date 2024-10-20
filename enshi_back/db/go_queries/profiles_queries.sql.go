// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: profiles_queries.sql

package db_repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const clearProfileByUserId = `-- name: ClearProfileByUserId :one
UPDATE public.profiles
SET bio='', avatar_url='', website_url=''
WHERE user_id=$1
RETURNING profile_id, user_id, bio, avatar_url, website_url
`

func (q *Queries) ClearProfileByUserId(ctx context.Context, userID pgtype.Int8) (Profile, error) {
	row := q.db.QueryRow(ctx, clearProfileByUserId, userID)
	var i Profile
	err := row.Scan(
		&i.ProfileID,
		&i.UserID,
		&i.Bio,
		&i.AvatarUrl,
		&i.WebsiteUrl,
	)
	return i, err
}

const createProfileForUser = `-- name: CreateProfileForUser :one
INSERT INTO public.profiles
(profile_id, user_id, bio, avatar_url, website_url)
VALUES($1, $2, '', '', '')
RETURNING profile_id, user_id, bio, avatar_url, website_url
`

type CreateProfileForUserParams struct {
	ProfileID int64       `json:"profile_id"`
	UserID    pgtype.Int8 `json:"user_id"`
}

func (q *Queries) CreateProfileForUser(ctx context.Context, arg CreateProfileForUserParams) (Profile, error) {
	row := q.db.QueryRow(ctx, createProfileForUser, arg.ProfileID, arg.UserID)
	var i Profile
	err := row.Scan(
		&i.ProfileID,
		&i.UserID,
		&i.Bio,
		&i.AvatarUrl,
		&i.WebsiteUrl,
	)
	return i, err
}

const deleteProfileByUserId = `-- name: DeleteProfileByUserId :exec
DELETE FROM public.profiles
WHERE user_id=$1
`

func (q *Queries) DeleteProfileByUserId(ctx context.Context, userID pgtype.Int8) error {
	_, err := q.db.Exec(ctx, deleteProfileByUserId, userID)
	return err
}

const getProfileByUserId = `-- name: GetProfileByUserId :one
SELECT profile_id, user_id, bio, avatar_url, website_url FROM public.profiles WHERE user_id = $1
`

func (q *Queries) GetProfileByUserId(ctx context.Context, userID pgtype.Int8) (Profile, error) {
	row := q.db.QueryRow(ctx, getProfileByUserId, userID)
	var i Profile
	err := row.Scan(
		&i.ProfileID,
		&i.UserID,
		&i.Bio,
		&i.AvatarUrl,
		&i.WebsiteUrl,
	)
	return i, err
}

const updateProfileByUserId = `-- name: UpdateProfileByUserId :one
UPDATE public.profiles
SET bio=$2, avatar_url=$3, website_url=$4
WHERE profile_id=$1
RETURNING profile_id, user_id, bio, avatar_url, website_url
`

type UpdateProfileByUserIdParams struct {
	ProfileID  int64       `json:"profile_id"`
	Bio        pgtype.Text `json:"bio"`
	AvatarUrl  pgtype.Text `json:"avatar_url"`
	WebsiteUrl pgtype.Text `json:"website_url"`
}

func (q *Queries) UpdateProfileByUserId(ctx context.Context, arg UpdateProfileByUserIdParams) (Profile, error) {
	row := q.db.QueryRow(ctx, updateProfileByUserId,
		arg.ProfileID,
		arg.Bio,
		arg.AvatarUrl,
		arg.WebsiteUrl,
	)
	var i Profile
	err := row.Scan(
		&i.ProfileID,
		&i.UserID,
		&i.Bio,
		&i.AvatarUrl,
		&i.WebsiteUrl,
	)
	return i, err
}
