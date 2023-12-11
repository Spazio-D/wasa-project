package database

import (
	"database/sql"
	"errors"
)

var query_is_following = "SELECT follower_id FROM Follow WHERE follower_id = ? AND followed_id = ?"

func (db *appdbimpl) IsFollowing(followerID int, followedID int) (bool, error) {
	var follower string
	err := db.c.QueryRow(query_is_following, followerID, followedID).Scan(&follower)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return follower != "", nil
}
