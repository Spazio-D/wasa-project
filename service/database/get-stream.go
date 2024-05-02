package database

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

		if err = db.c.QueryRow(query_like_check, post.ID, user.ID, userID).Scan(&post.LikeCheck); err != nil {
			return nil, err
		}

		rows2, err := db.c.Query(query_get_comments, post.ID, user.ID)
		if err != nil {
			return nil, err
		}

		var comments []Comment
		for rows2.Next() {
			var comment Comment
			err = rows2.Scan(&comment.ID, &comment.User.ID, &comment.PostID, &comment.OwnerID, &comment.Text, &comment.Timestamp)
			if err != nil {
				return nil, err
			}
			var user User
			user, err = db.GetUserByID(comment.User.ID)
			if err != nil {
				return nil, err
			}
			comment.User = user
			comments = append(comments, comment)
		}

		defer func() { err = rows2.Close() }()

		if rows2.Err() != nil {
			return posts, rows.Err()
		}

		post.User = user
		post.Comments = comments
		posts = append(posts, post)

	}

	defer func() { err = rows.Close() }()

	if rows.Err() != nil {
		return posts, rows.Err()
	}

	return posts, nil

}
