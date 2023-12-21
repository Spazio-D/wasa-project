package database

var query_get_comment_by_id = `SELECT id, user_id, owner_id, post_id, text, timestamp FROM comment WHERE id = ? AND owner_id = ? AND post_id = ?`

func (db *appdbimpl) GetCommentByID(id int, ownerID int, postID int) (Comment, error) {
	var comment Comment

	err := db.c.QueryRow(query_get_comment_by_id, id, ownerID, postID).Scan(&comment.ID, &comment.User.ID, &comment.OwnerID, &comment.PostID, &comment.Text, &comment.Timestamp)
	return comment, err
}
