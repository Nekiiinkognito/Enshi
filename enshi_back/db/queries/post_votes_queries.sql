-- name: CreatePostVote :one
INSERT INTO public.post_votes
(post_id, user_id, vote)
VALUES($1, $2, $3)
ON CONFLICT (user_id, post_id)
DO UPDATE SET
    vote = $3
RETURNING *;

-- name: DeletePostVote :exec
DELETE FROM public.post_votes
WHERE post_id=$1 AND user_id=$2;

-- name: UpdateVote :one
UPDATE public.post_votes
SET vote=$1
WHERE post_id=$2 AND user_id=$3
RETURNING *;

-- name: GetPostVote :one
SELECT vote
FROM public.post_votes p_v
WHERE p_v.user_id = $1 and p_v.post_id = $2;

-- name: GetPostVotes :one
SELECT count (*) FILTER (WHERE vote = TRUE) as upvotes,
count (*) FILTER (WHERE vote = FALSE) as downvotes
FROM public.post_votes
WHERE post_id = $1;
