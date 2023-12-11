package database

var query_delete_ban = `DELETE FROM Ban WHERE banner_id = ? AND banned_id = ?`

func (db *appdbimpl) DeleteBan(bannerID int, bannedID int) error {

	_, err := db.c.Exec(query_delete_ban, bannerID, bannedID)
	if err != nil {
		return err
	}
	return nil
}
