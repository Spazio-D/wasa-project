package database

import (
	"database/sql"
	"fmt"
	"os"
)

var query_max_user_id = `SELECT MAX(id) FROM User`
var query_insert_user = `INSERT INTO User (id, username) VALUES (?, ?)`

func (db *appdbimpl) CreateUser(user User) (User, error) {
	var id = sql.NullInt64{Int64: 0, Valid: false}

	err := db.c.QueryRow(query_max_user_id).Scan(&id)
	if !id.Valid {
		user.ID = 1
	} else if err != nil {
		return user, err
	} else {
		user.ID = int(id.Int64) + 1
	}

	_, err = db.c.Exec(query_insert_user, user.ID, user.Username)
	if err != nil {
		return user, err
	}

	path := "./users/" + user.Username + "_" + fmt.Sprint(user.ID) + "/posts"
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return user, err
	}

	return user, nil
}
