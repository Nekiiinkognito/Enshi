-- name: GetTagByTagId :one
SELECT tag_id, tag_name
FROM public.tags tags
where tags.tag_id = $1;

-- name: CreateTag :one
INSERT INTO public.tags
(tag_id, tag_name)
VALUES($1, $2)
RETURNING *;

-- name: DeleteTag :exec
DELETE FROM public.tags
WHERE tag_id = $1;

-- name: GetAllTags :many
SELECT tag_id, tag_name
FROM public.tags
ORDER BY tag_name ASC;