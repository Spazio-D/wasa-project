package database

var query_get_user = `SELECT id, username FROM User WHERE id = ?`
var query_get_followers = `SELECT id, username FROM User WHERE id IN (SELECT follower_id FROM Follow WHERE followed_id=?)`
var query_get_followed = `SELECT id, username FROM User WHERE id IN (SELECT followed_id FROM Follow WHERE follower_id=?)`
var query_get_post_count = `SELECT count(id) FROM Post WHERE user_id = ?`
var query_follow_check = `SELECT count(followed_id) FROM Follow WHERE followed_id = ? AND follower_id = ?`

func (db *appdbimpl) GetUserProfile(targetUserID int, askingUserID int) (Profile, error) {

	var profile Profile
	if err := db.c.QueryRow(query_get_user, targetUserID).Scan(&profile.User.ID, &profile.User.Username); err != nil {
		return profile, err
	}

	rows, err := db.c.Query(query_get_followers, targetUserID)
	if err != nil {
		return profile, err
	}

	for rows.Next() {
		if rows.Err() != nil {
			return profile, err
		}
		var follower User

		err = rows.Scan(&follower.ID, &follower.Username)
		if err != nil {
			return profile, err
		}

		profile.Followers = append(profile.Followers, follower)
	}

	err = rows.Close()
	if err != nil {
		return profile, err
	}

	rows, err = db.c.Query(query_get_followed, targetUserID)
	if err != nil {
		return profile, err
	}
	defer func() { err = rows.Close() }()

	for rows.Next() {
		if rows.Err() != nil {
			return profile, err
		}
		var follower User
		err = rows.Scan(&follower.ID, &follower.Username)
		if err != nil {
			return profile, err
		}

		profile.Followed = append(profile.Followed, follower)
	}

	err = rows.Close()
	if err != nil {
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

	profile.FollowersCount = len(profile.Followers)
	profile.FollowedCount = len(profile.Followed)
	return profile, nil
}
