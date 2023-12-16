package database

var query_DELETELIKE = `DELETE FROM Like WHERE user_id = ? AND owner_id = ? AND post_id = ?`

func (db *appdbimpl) DeleteLike(userID int, ownerID int, postID int) error {
	_, err := db.c.Exec(query_DELETELIKE, userID, ownerID, postID)
	if err != nil {
		return err
	}

	return nil
}
