package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
)

var query_max_user_id = `SELECT MAX(id) FROM User`
var query_insert_user = `INSERT INTO User (id, username) VALUES (?, ?)`

func (db *appdbimpl) CreateUser(user User) (User, error) {
	var id int

	err := db.c.QueryRow(query_max_user_id).Scan(&id)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		user.Id = 1
	} else if err != nil {
		return user, err
	} else {
		user.Id = id + 1
	}

	_, err = db.c.Exec(query_insert_user, user.Id, user.Username)
	if err != nil {
		return user, err
	}

	path := "./users/" + user.Username + "_" + fmt.Sprint(user.Id) + "/posts"
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return user, err
	}

	return user, nil
}
