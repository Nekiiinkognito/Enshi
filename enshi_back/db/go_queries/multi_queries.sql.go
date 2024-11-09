// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: multi_queries.sql

package db_repo

import (
	"context"
)

const getFavoriteBlogsByUserId = `-- name: GetFavoriteBlogsByUserId :many
SELECT blogs.blog_id, blogs.user_id, blogs.title, blogs.description, blogs.category_id, blogs.created_at
FROM favorites
JOIN blogs on blogs.blog_id = favorites.blog_id
WHERE favorites.user_id = $1
`

type GetFavoriteBlogsByUserIdRow struct {
	Blog Blog `json:"blog"`
}

func (q *Queries) GetFavoriteBlogsByUserId(ctx context.Context, userID int64) ([]GetFavoriteBlogsByUserIdRow, error) {
	rows, err := q.db.Query(ctx, getFavoriteBlogsByUserId, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFavoriteBlogsByUserIdRow
	for rows.Next() {
		var i GetFavoriteBlogsByUserIdRow
		if err := rows.Scan(
			&i.Blog.BlogID,
			&i.Blog.UserID,
			&i.Blog.Title,
			&i.Blog.Description,
			&i.Blog.CategoryID,
			&i.Blog.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
