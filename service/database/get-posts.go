package database

import (
	"database/sql"
	"errors"
)

var query_get_posts = `SELECT id, timestamp FROM Post WHERE user_id = ? ORDER BY timestamp DESC LIMIT ?, ?`
var query_get_likes_count = `SELECT COUNT(post_id) FROM Like WHERE post_id = ? AND owner_id = ?`
var query_get_comments_count = `SELECT COUNT(post_id) FROM Comment WHERE post_id = ? AND owner_id = ?`
var query_like_check = `SELECT COUNT(post_id) FROM Like WHERE post_id = ? AND owner_id = ? AND user_id = ?`

func (db *appdbimpl) GetPosts(userID int, owner User, offset int, limit int) ([]Post, error) {
	rows, err := db.c.Query(query_get_posts, owner.ID, offset, limit)
	if err != nil {
		return nil, err
	}

	var posts []Post

	for rows.Next() {
		var post Post
		post.User = owner
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

		var likeCheck int
		err = db.c.QueryRow(query_like_check, post.ID, post.User.ID, userID).Scan(&likeCheck)
		if errors.Is(err, sql.ErrNoRows) {
			post.LikeCheck = false
		} else if err != nil {
			return nil, err
		} else {
			post.LikeCheck = true
		}

		post.ImageUrl = post.GetPath()
		posts = append(posts, post)
	}

	defer func() { err = rows.Close() }()
	
	if rows.Err() != nil {
		return posts, rows.Err()
	}

	return posts, nil
}
