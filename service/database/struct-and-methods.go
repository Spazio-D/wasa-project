package database

import (
	"fmt"
	"time"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type Post struct {
	ID            int       `json:"id"`
	User          User      `json:"user"`
	ImageUrl      string    `json:"imageUrl"`
	LikesCount    int       `json:"likesCount"`
	CommentsCount int       `json:"commentsCount"`
	Comments      []Comment `json:"comments"`
	Timestamp     time.Time `json:"timestamp"`
	LikeCheck     bool      `json:"likeCheck"`
}

// Get the path of the post
func (post *Post) GetPath() string {
	return "./users/" + fmt.Sprint(post.User.ID) + "/posts/" + fmt.Sprint(post.ID) + ".jpeg"
}

type Profile struct {
	User           User   `json:"user"`
	FollowersCount int    `json:"followersCount"`
	FollowedCount  int    `json:"followedCount"`
	Followers      []User `json:"followers"`
	Followed       []User `json:"followed"`
	PostsCount     int    `json:"postsCount"`
	FollowCheck    bool   `json:"followCheck"`
	IsBanned       bool   `json:"isBanned"`
}

type Comment struct {
	ID        int       `json:"id"`
	PostID    int       `json:"postID"`
	OwnerID   int       `json:"ownerID"`
	User      User      `json:"user"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}
