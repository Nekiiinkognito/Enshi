// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db_repo

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Blog struct {
	BlogID      int64            `json:"blog_id"`
	UserID      int64            `json:"user_id"`
	Title       pgtype.Text      `json:"title"`
	Description pgtype.Text      `json:"description"`
	CategoryID  pgtype.Int4      `json:"category_id"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
}

type Bookmark struct {
	UserID       int64            `json:"user_id"`
	PostID       int64            `json:"post_id"`
	BookmarkedAt pgtype.Timestamp `json:"bookmarked_at"`
}

type Category struct {
	CategoryID   int32  `json:"category_id"`
	CategoryName string `json:"category_name"`
}

type Comment struct {
	CommentID int64            `json:"comment_id"`
	PostID    pgtype.Int8      `json:"post_id"`
	UserID    pgtype.Int8      `json:"user_id"`
	Content   pgtype.Text      `json:"content"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

type Favorite struct {
	UserID      int64            `json:"user_id"`
	BlogID      int64            `json:"blog_id"`
	FavoritedAt pgtype.Timestamp `json:"favorited_at"`
}

type Like struct {
	LikeID    int64            `json:"like_id"`
	UserID    pgtype.Int8      `json:"user_id"`
	CommentID pgtype.Int8      `json:"comment_id"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

type Post struct {
	PostID    int64            `json:"post_id"`
	BlogID    pgtype.Int8      `json:"blog_id"`
	UserID    int64            `json:"user_id"`
	Title     pgtype.Text      `json:"title"`
	Content   pgtype.Text      `json:"content"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

type PostTag struct {
	PostID int64 `json:"post_id"`
	TagID  int32 `json:"tag_id"`
}

type PostVote struct {
	PostID int64 `json:"post_id"`
	UserID int64 `json:"user_id"`
	Vote   bool  `json:"vote"`
}

type Profile struct {
	UserID     int64       `json:"user_id"`
	Bio        pgtype.Text `json:"bio"`
	AvatarUrl  pgtype.Text `json:"avatar_url"`
	WebsiteUrl pgtype.Text `json:"website_url"`
}

type Tag struct {
	TagID   int32  `json:"tag_id"`
	TagName string `json:"tag_name"`
}

type User struct {
	UserID    int64            `json:"user_id"`
	Username  string           `json:"username" validate:"required"`
	Email     string           `json:"email" validate:"required,email"`
	Password  string           `json:"password" validate:"required"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	IsAdmin   bool             `json:"is_admin"`
}
