-- name: CreateFavorite :one
INSERT INTO public.favorites
(user_id, blog_id, favorited_at)
VALUES($1, $2, CURRENT_TIMESTAMP)
RETURNING *;

-- name: DeleteFavorite :exec
DELETE FROM public.favorites
WHERE user_id=$1 AND blog_id=$2;