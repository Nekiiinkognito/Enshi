-- name: GetAllUsers :many
SELECT * FROM users;

-- name: GetUserById :one
SELECT * FROM users WHERE user_id = $1;

-- name: GetUserUsernameById :one
SELECT username FROM users WHERE user_id = $1;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1;

-- name: CreateUser :one
INSERT INTO public.users
(user_id, username, email, "password", created_at, is_admin)
VALUES($1, $2, $3, $4, CURRENT_TIMESTAMP, false)
RETURNING *;

-- name: UpdateUserPasswordHash :one
UPDATE public.users
SET "password"=$1
WHERE user_id=$2
RETURNING *;

-- name: DeleteUserById :exec
DELETE FROM public.users
WHERE user_id=$1;

-- name: DeleteUserByUsername :exec
DELETE FROM public.users
WHERE username=$1;

-- name: GetUserByEmailOrNickname :one
SELECT * FROM users WHERE username = $1 OR email = $2 LIMIT 1;