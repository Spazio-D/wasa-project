package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) IsLiked(postID int, ownerID int, userID int) (bool, error) {
	var ban int

	err := db.c.QueryRow(query_like_check, postID, ownerID, userID).Scan(&ban)

	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return ban != 0, err
}
