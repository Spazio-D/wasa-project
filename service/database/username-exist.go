package database

import (
	"database/sql"
	"errors"
)

var query_search_username = `SELECT username FROM User WHERE username = ?`

func (db *appdbimpl) UsernameExist(username string) (bool, error) {
	var result string
	err := db.c.QueryRow(query_search_username, username).Scan(&result)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return result != "", err
}
