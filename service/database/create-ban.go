package database

var query_create_ban = `INSERT INTO Ban (banner_id, banned_id) VALUES (?, ?)`

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

	_, err = db.c.Exec(query_create_ban, bannerID, bannedID)
	if err != nil {
		return err
	}
	return nil
}
