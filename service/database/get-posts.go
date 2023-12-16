package database

import (
	"database/sql"
	"errors"
)

var query_get_posts = `SELECT id, timestamp FROM Post WHERE user_id = ? ORDER BY timestamp DESC LIMIT ?, ?`
var query_get_likes_count = `SELECT COUNT(postID) FROM Like WHERE postID=? AND ownerID=?`
var query_get_comments_count = `SELECT COUNT(postID) FROM Comment WHERE postID=? AND ownerID=?`

func (db *appdbimpl) GetPosts(user User, offset int, limit int) ([]Post, error) {
	rows, err := db.c.Query(query_get_posts, user.ID, offset, limit)
	if err != nil {
		return nil, err
	}
	defer func() { err = rows.Close() }()

	var posts []Post

	for rows.Next() {
		var post Post
		post.User = user
		err = rows.Scan(&post.ID, &post.Timestamp)
		if err != nil {
			return nil, err
		}

		err = db.c.QueryRow(query_get_likes_count, post.ID, post.User.ID).Scan(&post.LikesCount)
		if errors.Is(err, sql.ErrNoRows) {
			post.LikesCount = 0
		} else if err != nil {
			return nil, err
		}

		err = db.c.QueryRow(query_get_comments_count, post.ID, post.User.ID).Scan(&post.CommentsCount)
		if errors.Is(err, sql.ErrNoRows) {
			post.LikesCount = 0
		} else if err != nil {
			return nil, err
		}

		post.ImageUrl = post.GetPath()
		posts = append(posts, post)
	}

	if rows.Err() != nil {
		return posts, err
	}

	return posts, nil
}
