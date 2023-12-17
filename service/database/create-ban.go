package database

var query_create_ban = `INSERT INTO Ban (banner_id, banned_id) VALUES (?, ?)`
var query_delete_comments = `DELETE FROM Comment WHERE user_id = ? AND owner_id = ?`
var query_delete_likes = `DELETE FROM Like WHERE user_id = ? AND owner_id = ?`

func (db *appdbimpl) CreateBan(bannerID int, bannedID int) error {

	follow, err := db.IsFollowing(bannerID, bannedID)
	if err != nil {
		return err
	}
	if follow {
		err = db.DeleteFollow(bannerID, bannedID)
		if err != nil {
			return err
		}
	}

	follow, err = db.IsFollowing(bannedID, bannerID)
	if err != nil {
		return err
	}
	if follow {
		err = db.DeleteFollow(bannedID, bannerID)
		if err != nil {
			return err
		}
	}

	_, err = db.c.Exec(query_delete_comments, bannedID, bannerID)
	if err != nil {
		return err
	}
	_, err = db.c.Exec(query_delete_comments, bannerID, bannedID)
	if err != nil {
		return err
	}

	_, err = db.c.Exec(query_delete_likes, bannedID, bannerID)
	if err != nil {
		return err
	}
	_, err = db.c.Exec(query_delete_likes, bannerID, bannedID)
	if err != nil {
		return err
	}

	_, err = db.c.Exec(query_create_ban, bannerID, bannedID)
	if err != nil {
		return err
	}
	return nil
}
