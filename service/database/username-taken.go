package database

import (
	"database/sql"
	"errors"
)

var query_searchUsername = `SELECT username FROM User WHERE username = ?`

func (db *appdbimpl) UsernameTaken(username string) (bool, error) {
	var result string
	err := db.c.QueryRow(query_searchUsername, username).Scan(&result)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	return result != "", err
}
