-- name: GetFavoriteBlogsInfosByUserId :many
SELECT sqlc.embed(blogs)
FROM favorites
JOIN blogs on blogs.blog_id = favorites.blog_id
WHERE favorites.user_id = $1;

