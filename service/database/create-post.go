package database

import (
	"Spazio-D/wasa-project/service/api/utils"
	"database/sql"
	"errors"
	"os"
)

var query_max_post_id = `SELECT MAX(id) FROM Post WHERE user_id = ?`
var query_create_post = `INSERT INTO Post (id, user_id,) VALUES (?, ?)`

func (db *appdbimpl) CreatePost(post Post, data []byte) (Post, error) {
	var id int
	err := db.c.QueryRow(query_max_post_id, post.User.ID).Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		post.ID = 1
	} else if err != nil {
		return post, err
	} else {
		post.ID = id + 1
	}

	path := post.GetPaht()
	err = utils.SaveAndResizeImage(path, data, os.ModePerm)
	if err != nil {
		return post, err
	}

	_, err = db.c.Exec(query_create_post, post.ID, post.User.ID)
	if err != nil {
		return post, err
	}

	user := post.User
	posts, err := db.GetPosts(user, 0, 1)
	if err != nil {
		return post, err
	}

	return posts[0], nil

}