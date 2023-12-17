package database

var query_create_like = `INSERT INTO Like (user_id, owner_id, post_id) VALUES (?, ?, ?)`

func (db *appdbimpl) CreateLike(userID int, ownerID int, postID int) error {
	_, err := db.c.Exec(query_create_like, userID, ownerID, postID)

	return err
}
