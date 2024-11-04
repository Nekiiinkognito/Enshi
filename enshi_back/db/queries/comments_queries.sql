-- name: CreateComment :one
INSERT INTO public."comments"
(comment_id, post_id, user_id, "content", created_at)
VALUES($1, $2, $3, $4, CURRENT_TIMESTAMP)
RETURNING *;

-- name: DeleteComment :exec
DELETE FROM public."comments"
WHERE comment_id = $1;

-- name: GetCommentsForPostDesc :many
SELECT comment_id, post_id, user_id, "content", created_at
FROM public."comments"
where public."comments".post_id = $1
order by created_at DESC
LIMIT 10 offset ($2 * 10);

-- name: GetCommentsForPostAsc :many
SELECT comment_id, post_id, user_id, "content", created_at
FROM public."comments"
where public."comments".post_id = $1
order by created_at ASC
LIMIT 10 offset ($2 * 10);

-- name: UpdateCommentByCommentId :one
UPDATE public."comments"
SET "content"=$2
WHERE comment_id=$1
RETURNING *;

-- name: GetCommentByUserId :one
SELECT comment_id, post_id, user_id, "content", created_at
FROM public."comments"
where public."comments".user_id = $1 and public."comments".post_id = $2;