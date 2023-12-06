package database

var query_change_username = `UPDATE User SET username = ? WHERE id = ?`

func (db *appdbimpl) ChangeUsername(username string, userId int) error {
	_, err := db.c.Exec(query_change_username, userId, username)
	return err
}

