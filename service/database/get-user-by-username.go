package database

var query_get_user_by_username = `SELECT id, username, FROM User WHERE username = ?`

func (db *appdbimpl) GetUserByUsername(username string) (User, error) {
	var user User
	err := db.c.QueryRow(query_get_user_by_username, username).Scan(&user.ID, &user.Username)
	return user, err
}
