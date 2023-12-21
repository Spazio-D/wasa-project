package database

import (
	"database/sql"
	"errors"
)

var get_user_followed = `SELECT id, username FROM User WHERE id IN (SELECT followed_id FROM Follow WHERE follower_id = ?)`
var query_get_stream = `SELECT User.id, User.username, Post.id, Post.timestamp FROM (` + get_user_followed + `) AS User INNER JOIN Post ON User.id = Post.user_id ORDER BY Post.timestamp DESC`

func (db *appdbimpl) GetStream(userID int) ([]Post, error) {
	var posts []Post

	rows, err := db.c.Query(query_get_stream, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user User
		var post Post

		err = rows.Scan(&user.ID, &user.Username, &post.ID, &post.Timestamp)
		if err != nil {
			return nil, err
		}

		if err := db.c.QueryRow(query_get_likes_count, post.ID, user.ID).Scan(&post.LikesCount); err != nil {
			return nil, err
		}

		if err := db.c.QueryRow(query_get_comments_count, post.ID, user.ID).Scan(&post.CommentsCount); err != nil {
			return nil, err
		}

		var likeCheck int
		err = db.c.QueryRow(query_like_check, post.ID, post.User.ID, userID).Scan(&likeCheck)
		if errors.Is(err, sql.ErrNoRows) {
			post.LikeCheck = false
		} else if err != nil {
			return nil, err
		} else {
			post.LikeCheck = true
		}

		post.User = user
		posts = append(posts, post)

	}

	defer func() { err = rows.Close() }()
	
	if rows.Err() != nil {
		return posts, rows.Err()
	}

	return posts, nil

}
