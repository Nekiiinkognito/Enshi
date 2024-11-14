-- name: GetPostsByPostId :one
SELECT *
FROM public.posts posts
where posts.post_id = $1;

-- name: GetPostsByUserId :many
SELECT *
FROM public.posts posts
where posts.user_id = $1;

-- name: GetPostsByBlogId :many
SELECT *
FROM public.posts posts 
where posts.blog_id = $1;

-- name: CreatePost :one
INSERT INTO public.posts
(post_id, blog_id, user_id, title, "content", created_at, updated_at)
VALUES($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING *;

-- name: UpdatePostByPostId :one
UPDATE public.posts
SET blog_id=$1, title=$2, "content"=$3, updated_at=CURRENT_TIMESTAMP
WHERE post_id = $4
RETURNING *;

-- name: DeletePostByPostId :exec
DELETE FROM public.posts
WHERE post_id=$1;