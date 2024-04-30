package database

var query_get_posts = `SELECT id, timestamp FROM Post WHERE user_id = ? ORDER BY timestamp DESC LIMIT ?, ?`
var query_get_likes_count = `SELECT COUNT(post_id) FROM Like WHERE post_id = ? AND owner_id = ?`
var query_get_comments_count = `SELECT COUNT(post_id) FROM Comment WHERE post_id = ? AND owner_id = ?`
var query_get_comments = `SELECT * FROM Comment WHERE post_id = ? AND owner_id = ?`
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
		if err != nil {
			return nil, err
		}

		err = db.c.QueryRow(query_get_comments_count, post.ID, post.User.ID).Scan(&post.CommentsCount)
		if err != nil {
			return nil, err
		}

		err = db.c.QueryRow(query_like_check, post.ID, post.User.ID, userID).Scan(&post.LikeCheck)
		if err != nil {
			return nil, err
		}

		rows2, err := db.c.Query(query_get_comments, post.ID, owner.ID)
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

		post.ImageUrl = post.GetPath()
		post.Comments = comments
		posts = append(posts, post)

		defer func() { err = rows2.Close() }()
	}

	defer func() { err = rows.Close() }()

	if rows.Err() != nil {
		return posts, rows.Err()
	}

	return posts, nil
}
