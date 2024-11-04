-- name: CreateLike :one
INSERT INTO public.likes
(like_id, user_id, comment_id, created_at)
VALUES($1, $2, $3, CURRENT_TIMESTAMP)
RETURNING *;

-- name: DeleteLikeByUserCommentId :exec
DELETE FROM public.likes
WHERE user_id = $1 and comment_id = $2;

-- name: GetLikesForComment :one
SELECT count(*)
FROM public.likes
WHERE comment_id = $1;

-- name: IsUserLikedComment :one
SELECT count(*)
FROM public.likes
WHERE user_id = $1 and comment_id = $2;

