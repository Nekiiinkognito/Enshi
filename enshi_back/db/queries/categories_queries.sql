-- name: CreateCategory :one
INSERT INTO public.categories
(category_id, category_name)
VALUES($1, $2)
RETURNING *;

-- name: GetAllCategories :many
SELECT * FROM public.categories;

-- name: GetCategoryByName :one
SELECT * FROM public.categories WHERE category_name = $1;

-- name: DeleteCategoryById :exec
DELETE FROM public.categories
WHERE category_id=$1;

