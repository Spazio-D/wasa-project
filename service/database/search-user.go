package database

var query_get_users = `SELECT id, username FROM User WHERE username regexp ? ORDER BY username`

func (db *appdbimpl) SearchUsers(userID int, username string) ([]User, error) {
	var users []User

	rows, err := db.c.Query(query_get_users, "^"+username)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Username)
		if err != nil {
			return nil, err
		}

		banCheck, err := db.IsBanned(user.ID, userID)
		if err != nil {
			return nil, err
		}

		if !banCheck {
			users = append(users, user)
		}
	}

	defer func() { err = rows.Close() }()

	if rows.Err() != nil {
		return users, rows.Err()
	}

	return users, err
}
