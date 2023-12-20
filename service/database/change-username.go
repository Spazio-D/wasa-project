package database

var query_change_username = `UPDATE User SET username = ? WHERE id = ?`

func (db *appdbimpl) ChangeUsername(username string, userID int) error {
	_, err := db.c.Exec(query_change_username, username, userID)
	return err
}
