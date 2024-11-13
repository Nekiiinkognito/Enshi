package checkRole

import (
	"context"
	db_repo "enshi/db/go_queries"
	"enshi/db_connection"
)

func IsOwnerOfThePost(userId int64, postId int64) (bool, error) {
	post, err :=
		db_repo.New(db_connection.Dbx).
			GetPostsByPostId(context.Background(), postId)

	if err != nil {
		return false, err
	}

	if post.UserID != userId {
		return false, nil
	}

	return true, nil
}
