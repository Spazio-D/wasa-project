package database

var query_delete_like = `DELETE FROM Like WHERE user_id = ? AND owner_id = ? AND post_id = ?`

func (db *appdbimpl) DeleteLike(userID int, ownerID int, postID int) error {
	_, err := db.c.Exec(query_delete_like, userID, ownerID, postID)
	if err != nil {
		return err
	}

	return nil
}
