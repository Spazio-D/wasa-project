package database

import (
	"fmt"
	"time"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type Post struct {
	Id            int       `json:"id"`
	User          User      `json:"user"`
	ImageUrl      string    `json:"imageUrl"`
	LikesCount    int       `json:"likesCount"`
	CommentsCount int       `json:"commentsCount"`
	Timestamp     time.Time `json:"timestamp"`
}

// Get the path of the post
func (post *Post) GetPaht() string {
	return "./users/" + post.User.Username + "_" + fmt.Sprint(post.User.Id) + "/posts/" + fmt.Sprint(post.Id) + ".jpeg"
}
