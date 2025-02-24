// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users_queries.sql

package db_repo

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO public.users
(user_id, username, email, "password", created_at, is_admin)
VALUES($1, $2, $3, $4, CURRENT_TIMESTAMP, false)
RETURNING user_id, username, email, password, created_at, is_admin
`

type CreateUserParams struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.UserID,
		arg.Username,
		arg.Email,
		arg.Password,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.IsAdmin,
	)
	return i, err
}

const deleteUserById = `-- name: DeleteUserById :exec
DELETE FROM public.users
WHERE user_id=$1
`

func (q *Queries) DeleteUserById(ctx context.Context, userID int64) error {
	_, err := q.db.Exec(ctx, deleteUserById, userID)
	return err
}

const deleteUserByUsername = `-- name: DeleteUserByUsername :exec
DELETE FROM public.users
WHERE username=$1
`

func (q *Queries) DeleteUserByUsername(ctx context.Context, username string) error {
	_, err := q.db.Exec(ctx, deleteUserByUsername, username)
	return err
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT user_id, username, email, password, created_at, is_admin FROM users
`

func (q *Queries) GetAllUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.UserID,
			&i.Username,
			&i.Email,
			&i.Password,
			&i.CreatedAt,
			&i.IsAdmin,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByEmailOrNickname = `-- name: GetUserByEmailOrNickname :one
SELECT user_id, username, email, password, created_at, is_admin FROM users WHERE username = $1 OR email = $2 LIMIT 1
`

type GetUserByEmailOrNicknameParams struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

func (q *Queries) GetUserByEmailOrNickname(ctx context.Context, arg GetUserByEmailOrNicknameParams) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmailOrNickname, arg.Username, arg.Email)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.IsAdmin,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT user_id, username, email, password, created_at, is_admin FROM users WHERE user_id = $1
`

func (q *Queries) GetUserById(ctx context.Context, userID int64) (User, error) {
	row := q.db.QueryRow(ctx, getUserById, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.IsAdmin,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT user_id, username, email, password, created_at, is_admin FROM users WHERE username = $1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.IsAdmin,
	)
	return i, err
}

const getUserUsernameById = `-- name: GetUserUsernameById :one
SELECT username FROM users WHERE user_id = $1
`

func (q *Queries) GetUserUsernameById(ctx context.Context, userID int64) (string, error) {
	row := q.db.QueryRow(ctx, getUserUsernameById, userID)
	var username string
	err := row.Scan(&username)
	return username, err
}

const updateUserPasswordHash = `-- name: UpdateUserPasswordHash :one
UPDATE public.users
SET "password"=$1
WHERE user_id=$2
RETURNING user_id, username, email, password, created_at, is_admin
`

type UpdateUserPasswordHashParams struct {
	Password string `json:"password" validate:"required"`
	UserID   int64  `json:"user_id"`
}

func (q *Queries) UpdateUserPasswordHash(ctx context.Context, arg UpdateUserPasswordHashParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUserPasswordHash, arg.Password, arg.UserID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.IsAdmin,
	)
	return i, err
}
