package api

import (
	"Spazio-D/wasa-project/service/api/utils"
	"Spazio-D/wasa-project/service/database"
	"regexp"
	"time"
)

// USER STRUCT AND METHODS
type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

// Check if the username is valid
func (user *User) IsValid() bool {
	return regexp.MustCompile(`^\w{3,16}$`).MatchString(user.Username)
}

// Convert the user from api struct to database struct
func (user *User) DatabaseConversion() database.User {
	return database.User{
		Id:       user.Id,
		Username: user.Username,
	}
}

// Convert the user from database struct to api struct
func (user *User) ApiConversion(dbUser database.User) {
	user.Id = dbUser.Id
	user.Username = dbUser.Username
}

// AUTORIZATION USER STRUCT
type AuthUser struct {
	User  User `json:"user"`
	Token int  `json:"token"`
}

// POST STRUCT AND METHODS
type Post struct {
	Id            int       `json:"id"`
	User          User      `json:"user"`
	Image         string    `json:"image"`
	LikesCount    int       `json:"likesCount"`
	CommentsCount int       `json:"commentsCount"`
	Timestamp     time.Time `json:"timestamp"`
}

// Convert the post from api struct to database struct
func (post *Post) DatabaseConversion() database.Post {
	return database.Post{
		Id:            post.Id,
		User:          database.User{Id: post.User.Id, Username: post.User.Username},
		LikesCount:    post.LikesCount,
		CommentsCount: post.CommentsCount,
		Timestamp:     post.Timestamp,
	}
}

// Convert the post from database struct to api struct
func (post *Post) ApiConversion(dbPost database.Post) error {
	image, err := utils.Base64(dbPost.GetPaht())
	if err != nil {
		return err
	}

	var user User
	user.ApiConversion(dbPost.User)

	post.Id = dbPost.Id
	post.User = user
	post.Image = image
	post.LikesCount = dbPost.LikesCount
	post.CommentsCount = dbPost.CommentsCount
	post.Timestamp = dbPost.Timestamp

	return nil
}
