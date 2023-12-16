package database

var get_user_followed = `SELECT id, username FROM User WHERE id IN (SELECT followed_id FROM Follow WHERE follower_id = ?)`
var query_get_stream = `SELECT User.id, User.username, Post.id, Post.timestamp FROM (` + get_user_followed + `) AS User INNER JOIN Post ON User.id = Post.id ORDER BY Post.timestamp DESC`

func (db *appdbimpl) GetStream(userID int) ([]Post, error) {
	var posts []Post

	rows, err := db.c.Query(query_get_stream, userID)
	if err != nil {
		return nil, err
	}
	defer func() { err = rows.Close() }()

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

		post.User = user
		posts = append(posts, post)

	}

	return posts, nil

}
