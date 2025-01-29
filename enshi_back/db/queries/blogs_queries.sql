-- name: CreateBlogByUserId :one
INSERT INTO public.blogs
(blog_id, user_id, title, description, category_id, created_at)
VALUES($1, $2, $3, $4, $5, CURRENT_TIMESTAMP)
RETURNING *;

-- name: UpdateBlogInfoByBlogId :one
UPDATE public.blogs
SET title=$1, description=$2, category_id=$3
WHERE blog_id=$4
RETURNING *;

-- name: GetBlogsByUserId :many
SELECT *
FROM public.blogs
WHERE user_id = $1;

-- name: GetBlogByBlogId :one
SELECT *
FROM public.blogs
WHERE blog_id = $1;

-- name: DeleteBlogByBlogId :exec
DELETE FROM public.blogs
WHERE blog_id=$1;