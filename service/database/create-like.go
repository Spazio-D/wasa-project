package database

var query_CREATELIKE = `INSERT INTO Like (user_id, owner_id, post_id) VALUES (?, ?, ?)`

func (db *appdbimpl) CreateLike(userID int, ownerID int, postID int) error {
	_, err := db.c.Exec(query_CREATELIKE, userID, ownerID, postID)

	return err
}
