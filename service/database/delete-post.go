package database

import (
	"fmt"
	"os"
)

var query_delete_post = `DELETE FROM Post WHERE id = ? AND user_id = ?`

func (db *appdbimpl) DeletePost(postID int, userID int) error {
	_, err := db.c.Exec(query_delete_post, postID, userID)
	if err != nil {
		return err
	}

	var user User
	user, err = db.GetUserByID(userID)
	if err != nil {
		return err
	}

	err = os.Remove("./users/" + user.Username + "_" + fmt.Sprint(userID) + "/posts/" + fmt.Sprint(postID) + ".jpeg")
	if err != nil {
		return err
	}

	return nil
}
