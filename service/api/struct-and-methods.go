package api

import (
	"Spazio-D/wasa-project/service/api/utils"
	"Spazio-D/wasa-project/service/database"
	"regexp"
	"time"
)

// USER STRUCT AND METHODS
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

// Check if the username is valid
func (user *User) IsValid() bool {
	return regexp.MustCompile(`^\w{3,16}$`).MatchString(user.Username)
}

// Convert the user from api struct to database struct
func (user *User) DatabaseConversion() database.User {
	return database.User{
		ID:       user.ID,
		Username: user.Username,
	}
}

// Convert the user from database struct to api struct
func (user *User) ApiConversion(dbUser database.User) {
	user.ID = dbUser.ID
	user.Username = dbUser.Username
}

// AUTORIZATION USER STRUCT
type AuthUser struct {
	User  User `json:"user"`
	Token int  `json:"token"`
}

// POST STRUCT AND METHODS
type Post struct {
	ID            int       `json:"id"`
	User          User      `json:"user"`
	Image         string    `json:"image"`
	LikesCount    int       `json:"likesCount"`
	CommentsCount int       `json:"commentsCount"`
	Comments	  []Comment `json:"comments"`
	Timestamp     time.Time `json:"timestamp"`
	LikeCheck     bool      `json:"likeCheck"`
}

// Convert the post from api struct to database struct
func (post *Post) DatabaseConversion() database.Post {
	return database.Post{
		ID:            post.ID,
		User:          database.User{ID: post.User.ID, Username: post.User.Username},
		LikesCount:    post.LikesCount,
		CommentsCount: post.CommentsCount,
		Timestamp:     post.Timestamp,
		LikeCheck:     post.LikeCheck,
	}
}

// Convert the post from database struct to api struct
func (post *Post) ApiConversion(dbPost database.Post) error {
	image, err := utils.Base64(dbPost.GetPath())
	if err != nil {
		return err
	}

	var user User
	user.ApiConversion(dbPost.User)

	post.ID = dbPost.ID
	post.User = user
	post.Image = image
	post.LikesCount = dbPost.LikesCount
	post.CommentsCount = dbPost.CommentsCount
	post.Comments = make([]Comment, len(dbPost.Comments))
	for i, dbComment := range dbPost.Comments {
		var comment Comment
		comment.ApiConversion(dbComment)
		post.Comments[i] = comment
	}
	post.Timestamp = dbPost.Timestamp
	post.LikeCheck = dbPost.LikeCheck

	return nil
}

// PROFILE STRUCT AND METHODS
type Profile struct {
	User          User   `json:"user"`
	FollowerCount int    `json:"followersCount"`
	FollowedCount int    `json:"followedCount"`
	Followers     []User `json:"followers"`
	Followed      []User `json:"followed"`
	PostsCount    int    `json:"postsCount"`
	FollowCheck   bool   `json:"followCheck"`
	IsBanned      bool   `json:"isBanned"`
}

func (profile *Profile) ApiConversion(dbProfile database.Profile) {
	var user User
	user.ApiConversion(dbProfile.User)

	profile.User = user
	profile.FollowerCount = dbProfile.FollowersCount
	profile.Followers = make([]User, len(dbProfile.Followers))
	for i, dbUser := range dbProfile.Followers {
		var apiUser User
		apiUser.ApiConversion(dbUser)
		profile.Followers[i] = apiUser
	}
	profile.FollowedCount = dbProfile.FollowedCount
	profile.Followed = make([]User, len(dbProfile.Followed))
	for i, dbUser := range dbProfile.Followed {
		var apiUser User
		apiUser.ApiConversion(dbUser)
		profile.Followed[i] = apiUser
	}
	profile.PostsCount = dbProfile.PostsCount
	profile.FollowCheck = dbProfile.FollowCheck
	profile.IsBanned = dbProfile.IsBanned

}

// COMMENT STRUCT AND METHODS
type Comment struct {
	ID        int       `json:"id"`
	PostID    int       `json:"postID"`
	OwnerID   int       `json:"ownerID"`
	User      User      `json:"user"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}

func (comment *Comment) ApiConversion(dbComment database.Comment) {
	var user User
	user.ApiConversion(dbComment.User)

	comment.ID = dbComment.ID
	comment.PostID = dbComment.PostID
	comment.OwnerID = dbComment.OwnerID
	comment.User = user
	comment.Text = dbComment.Text
	comment.Timestamp = dbComment.Timestamp
}

func (comment *Comment) IsValid() bool {
	return regexp.MustCompile(`^.{1,999}$`).MatchString(comment.Text)
}
