package database

var query_DELETECOMMENT = `DELETE FROM Comment WHERE id = ? AND owner_id = ? AND post_id = ?`

func (db *appdbimpl) DeleteComment(userID int, postID int, commentID int) error {
	_, err := db.c.Exec(query_DELETECOMMENT, commentID, userID, postID)
	if err != nil {
		return err
	}

	return nil
}