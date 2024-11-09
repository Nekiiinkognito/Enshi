-- name: CreateProfileForUser :one
INSERT INTO public.profiles
(user_id, bio, avatar_url, website_url)
VALUES($1, '', '', '')
RETURNING *;

-- name: ClearProfileByUserId :one
UPDATE public.profiles
SET bio='', avatar_url='', website_url=''
WHERE user_id=$1
RETURNING *;

-- name: DeleteProfileByUserId :exec
DELETE FROM public.profiles
WHERE user_id=$1;

-- name: GetProfileByUserId :one
SELECT * FROM public.profiles WHERE user_id = $1;

-- name: UpdateProfileByUserId :one
UPDATE public.profiles
SET bio=$2, avatar_url=$3, website_url=$4
WHERE user_id=$1
RETURNING *;

