package database

var query_create_follow = "INSERT INTO Follow (follower_id, followed_id) VALUES (?, ?)"

func (db *appdbimpl) CreateFollow(followerID int, followedID int) error {
	_, err := db.c.Exec(query_create_follow, followerID, followedID)
	return err
}
