package database

import (
	"database/sql"
	"errors"
)

var query_is_banned = `SELECT count(banner_id) FROM Ban WHERE (banner_id = ? AND banned_id = ?)`

func (db *appdbimpl) IsBanned(bannerID int, bannedID int) (bool, error) {
	var ban int

	err := db.c.QueryRow(query_is_banned, bannerID, bannedID).Scan(&ban)

	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return ban != 0, err
}
