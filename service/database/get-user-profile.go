package database

var query_get_user = `SELECT id, username FROM User WHERE id = ?`
var query_get_followers_count = `SELECT count(followed_id) FROM Follow WHERE follower_id = ?`
var query_get_followed_count = `SELECT count(follower_id) FROM Follow WHERE followed_id = ?`
var query_get_post_count = `SELECT count(id) FROM Post WHERE user_id = ?`
var query_follow_check = `SELECT count(followed_id) FROM Follow WHERE followed_id = ? AND follower_id = ?`

func (db *appdbimpl) GetUserProfile(targetUserID int, askingUserID int) (Profile, error) {

	var profile Profile
	if err := db.c.QueryRow(query_get_user, targetUserID).Scan(&profile.User.ID, &profile.User.Username); err != nil {
		return profile, err
	}

	if err := db.c.QueryRow(query_get_followers_count, targetUserID).Scan(&profile.FollowersCount); err != nil {
		return profile, err
	}

	if err := db.c.QueryRow(query_get_followed_count, targetUserID).Scan(&profile.FollowedCount); err != nil {
		return profile, err
	}

	if err := db.c.QueryRow(query_get_post_count, targetUserID).Scan(&profile.PostsCount); err != nil {
		return profile, err
	}

	var followCheck int
	if err := db.c.QueryRow(query_follow_check, targetUserID, askingUserID).Scan(&followCheck); err != nil {
		return profile, err
	}
	if followCheck == 0 {
		profile.FollowCheck = false
	} else {
		profile.FollowCheck = true
	}

	return profile, nil

}
