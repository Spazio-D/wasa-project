package database

import (
	"database/sql"
	"time"
)

var query_max_comment_id = `SELECT MAX(id) FROM Comment WHERE owner_id = ? AND post_id = ?`
var query_create_comment = `INSERT INTO Comment (id, user_id, owner_id, post_id, text) VALUES (?, ?, ?, ?, ?)`

func (db *appdbimpl) CreateComment(userID int, ownerID int, postID int, text string) (Comment, error) {
	var id = sql.NullInt64{Int64: 0, Valid: false}
	var comment Comment

	err := db.c.QueryRow(query_max_comment_id, ownerID, postID).Scan(&id)
	if !id.Valid {
		comment.ID = 1
	} else if err != nil {
		return comment, err
	} else {
		comment.ID = int(id.Int64) + 1
	}

	_, err = db.c.Exec(query_create_comment, comment.ID, userID, ownerID, postID, text)
	if err != nil {
		return comment, err
	}

	user, err := db.GetUserByID(userID)
	if err != nil {
		return comment, err
	}

	comment.User = user
	comment.PostID = postID
	comment.OwnerID = ownerID
	comment.Text = text
	comment.Timestamp = time.Now()

	return comment, nil

}
