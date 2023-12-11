/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type appdbimpl struct {
	c   *sql.DB
	ctx context.Context
}

// AppDatabase is the high level interface for the DB
type AppDatabase interface {

	//Check if the username is in the database
	UsernameExist(username string) (bool, error)

	//Create a new user and return it
	CreateUser(user User) (User, error)

	//Get a user by its username
	GetUserByUsername(username string) (User, error)

	//Change the username of a user
	ChangeUsername(username string, userID int) error

	//Get a user by its id
	GetUserByID(id int) (User, error)

	//Create a new post and return it
	CreatePost(post Post, data []byte) (Post, error)

	//Return true if the followerID user is following the followedID user
	IsFollowing(followerID int, followedID int) (bool, error)

	//Return true if the bannerID user has banned the bannedID user
	IsBanned(bannerID int, bannedID int) (bool, error)

	//Create a follow relationship between the users
	CreateFollow(followerID int, followedID int) error

	//Delete a follow relationship between the users
	DeleteFollow(followerID int, followedID int) error

	//Create a ban relationship between two users
	CreateBan(bannerID int, bannedID int) error

	//Delete a ban relationship between two users
	DeleteBan(bannerID int, bannedID int) error

	//Get the profile of a user
	GetUserProfile(targetUserID int, askingUserID int) (Profile, error)

	Ping() error
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='example_table';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE example_table (id INTEGER NOT NULL PRIMARY KEY, name TEXT);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c:   db,
		ctx: context.Background(),
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
