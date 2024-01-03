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
	"database/sql"
	"errors"
	"fmt"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	GetStream(userid string) ([]Photo, error)

	CreateUser(userid string) error

	SetMyUsername(userid string, newUserid string) error

	FollowUser(followed string, follower string) error

	UnfollowUser(exFollowed string, exFollower string) error

	BanUser(banned string, banner string) error

	UnbanUser(exBanned string, exBanner string) error

	CreatePhoto(Photo) (int, error)

	DeletePhoto(photoid int, userid string) error

	LikePhoto(photoid int, userid string) error

	UnlikePhoto(photoid int, userid string) error

	CommentPhoto(photoid int, userid string, commentText string) (int, error)

	UncommentPhoto(photoid int, userid string, commentid int) error

	// Other methods

	GetFollowing(userid string) ([]string, error)
	GetFollowers(userid string) ([]string, error)
	BanCheck(banned string, banner string) (bool, error)
	GetPhotoList(userid string) ([]Photo, error)
	GetPhotoComments(photoid int) ([]Comment, error)
	CheckUser(userid string) (bool, error)

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Enable foreign keys for database
	_, errFK := db.Exec("PRAGMA foreign_keys = ON")
	if errFK != nil {
		return nil, fmt.Errorf("error in setting pragmas: %w", errFK)
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='users';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {

		wasaDatabase := [7]string{
			`CREATE TABLE users (
				userid VARCHAR(16) PRIMARY KEY 
				);`,
			`CREATE TABLE photos (
				photoid INTEGER PRIMARY KEY,
				userid VARCHAR(16) NOT NULL, 
				date DATETIME NOT NULL,
				FOREIGN KEY(userid) REFERENCES users (userid) ON DELETE CASCADE
				);`,
			`CREATE TABLE comments (
				commentid INTEGER PRIMARY KEY,
				userid VARCHAR(16) NOT NULL, 
				photoid INTEGER NOT NULL,
				commentText TEXT NOT NULL,
				FOREIGN KEY(photoid) REFERENCES photos (photoid),
				FOREIGN KEY(userid) REFERENCES photos (userid)
				);`,
			`CREATE TABLE followers (
				followerid VARCHAR(16) NOT NULL,
				followedid VARCHAR(16) NOT NULL,
				PRIMARY KEY (followedid, followerid), 
				FOREIGN KEY(followerid) REFERENCES users (userid) ON DELETE CASCADE,
				FOREIGN KEY(followedid) REFERENCES users (userid) ON DELETE CASCADE  
				);`,
			`CREATE TABLE likes (
				photoid INTEGER NOT NULL,
				userid VARCHAR(16) NOT NULL,
				PRIMARY KEY (userid, photoid),
				FOREIGN KEY(userid) REFERENCES users (userid) ON DELETE CASCADE,
				FOREIGN KEY(photoid) REFERENCES photos (photoid) ON DELETE CASCADE
				);`,
			`CREATE TABLE banned (
				bannedid VARCHAR(16) NOT NULL,
				bannerid VARCHAR(16) NOT NULL,
				PRIMARY KEY (bannerid, bannedid),
				FOREIGN KEY(bannedid) REFERENCES users (userid) ON DELETE CASCADE,
				FOREIGN KEY(bannerid) REFERENCES users (userid) ON DELETE CASCADE
				);`,
		}
		for i := 0; i < len(wasaDatabase); i++ {
			sqlStmt := wasaDatabase[i]
			_, err = db.Exec(sqlStmt)
			if err != nil {
				return nil, fmt.Errorf("error creating database structure: %w", err)
			}
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
