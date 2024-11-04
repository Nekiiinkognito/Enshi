-- name: GetBookmarksByUserId :many
SELECT *
FROM public.bookmarks bookmarks
where bookmarks.user_id = $1;

-- name: GetCountOfBookmarksByPostId :one
SELECT COUNT(*)
FROM public.bookmarks bookmarks
where bookmarks.post_id = $1;

-- name: GetBookmarkTimestamp :one
SELECT bookmarked_at
FROM public.bookmarks bookmarks
where bookmarks.post_id = $1 and bookmarks.user_id = $2;

-- name: CreateBookmark :one
INSERT INTO public.bookmarks
(user_id, post_id, bookmarked_at)
VALUES($1, $2, CURRENT_TIMESTAMP)
RETURNING *;

-- name: DeleteBookmark :exec
DELETE FROM public.bookmarks
WHERE user_id=$1 AND post_id=$2;