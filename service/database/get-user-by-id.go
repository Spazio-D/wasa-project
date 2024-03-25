package database

var query_get_user_by_id = `SELECT id, username FROM User WHERE id = ?`

func (db *appdbimpl) GetUserByID(id int) (User, error) {

	var user User
	err := db.c.QueryRow(query_get_user_by_id, id).Scan(&user.ID, &user.Username)

	return user, err
}
