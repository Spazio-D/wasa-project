package database

var query_is_banned = `SELECT banner_id, banned_id FROM Ban WHERE (banner_id = ? AND banned_id = ?)`

func (db *appdbimpl) IsBanned(bannerID int, bannedID int) (bool, error) {
	var ban string
	err := db.c.QueryRow(query_is_banned, bannerID, bannedID).Scan(&ban)

	if err != nil {
		return false, err
	}

	return ban != "", err
}
