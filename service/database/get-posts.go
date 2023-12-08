package database

//RICORDA DI CHIUDURE I *sql.Rows E METTERE I rows.Err()

var query_get_posts = `SELECT id, timestamp FROM Post WHERE user_id = ? ORDER BY timestamp DESC LIMIT ?, ?`
var query_get_likes_count = `SELECT COUNT(postID) FROM Like WHERE postID=? AND ownerID=?`
var query_get_comments_count = `SELECT COUNT(postID) FROM Comment WHERE postID=? AND ownerID=?`

func (db *appdbimpl) GetPosts(user User, offset int, limit int) ([]Post, error) {
	rows, err := db.c.Query(query_get_posts, user.Id, offset, limit)
	if err != nil {
		return nil, err
	}
	defer func() { err = rows.Close() }()

	var posts []Post

	for rows.Next() {
		var post Post
		post.User = user
		err = rows.Scan(&post.Id, &post.Timestamp)
		if err != nil {
			return nil, err
		}

		err = db.c.QueryRow(query_get_likes_count, post.Id, post.User.Id).Scan(&post.LikesCount)
		if err != nil {
			return nil, err
		}

		err = db.c.QueryRow(query_get_comments_count, post.Id, post.User.Id).Scan(&post.CommentsCount)
		if err != nil {
			return nil, err
		}

		post.ImageUrl = post.GetPaht()
		posts = append(posts, post)
	}
	return posts, nil
}
