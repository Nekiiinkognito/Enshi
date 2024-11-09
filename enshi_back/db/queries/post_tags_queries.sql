-- name: GetAllTagsForPost :many
SELECT sqlc.embed(tags)
from public.tags tags
JOIN public.post_tags post_tags on post_tags.tag_id = tags.tag_id
JOIN public.posts posts on posts.post_id = post_tags.post_id;

-- name: CreatePostTagRelation :one
INSERT INTO public.post_tags
(post_id, tag_id)
VALUES($1, $2)
RETURNING *; 

-- name: DeletePostTagRelation :exec
DELETE FROM public.post_tags
WHERE post_id = $1 AND tag_id = $2;