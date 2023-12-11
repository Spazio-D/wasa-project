package database

var query_delete_follow = "DELETE FROM Follow WHERE follower_id = ? AND followed_id = ?"

func (db *appdbimpl) DeleteFollow(followerID int, followedID int) error {
	_, err := db.c.Exec(query_delete_follow, followerID, followedID)
	return err
}
