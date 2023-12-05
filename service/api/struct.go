package api

import (
	"Spazio-D/wasa-project/service/database"
	"regexp"
)

// USER STRUCT AND METHODS
type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

// Check if the username is valid
func (user *User) isValid() bool {
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
