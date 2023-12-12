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

	//GetStream
	GetStream(UserId) ([]Photo, error)

	//CreateUser
	CreateUser(Username) (int, error)

	//SetMyUsername
	SetMyUsername(UserId, Username) error

	//FollowUser
	FollowUser(newFollowed UserId, newFollower UserId) error

	//UnfollowUser
	UnfollowUser(oldFollowed UserId, oldFollower UserId) error

	//BanUser
	BanUser(newBanned UserId, newBanner UserId) error

	//UnbanUser
	UnbanUser(oldBanned UserId, oldBanner UserId) error

	//CreatePhoto
	CreatePhoto(Photo) (int, error)

	//DeletePhoto
	DeletePhoto(Photo) error

	//LikePhoto
	LikePhoto(PhotoId, UserId) error

	//UnlikePhoto
	UnlikePhoto(PhotoId, UserId) error

	//CommentPhoto
	CommentPhoto(PhotoId, UserId, CommentText) (int, error)

	//UncommentPhoto
	UncommentPhoto(PhotoId, UserId, CommentId) error

	//Other methods

	GetFollowing(UserId) ([]UserId, error)
	GetFollowers(UserId) ([]UserId, error)
	BanCheck(UserId, UserId) (bool, error)
	GetPhotoList(UserId) ([]Photo, error)
	GetPhotoComments(PhotoId) ([]Comment, error)
	GetUsername(UserId) (string, error)

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

	//Enable foreign keys for database (https://www.sqlite.org/foreignkeys.html)
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
				userid INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
				username VARCHAR(16) NOT NULL
				);`,
			//AUTOINCREMENT: each user is guaranteed to have userid never used before by the same table in the same database
			//VARCHAR() instead of TEXT because TEXT is better suited for large amount of data
			`CREATE TABLE photos (
				photoid INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
				userid INTEGER NOT NULL, 
				date DATETIME NOT NULL,
				FOREIGN KEY(userid) REFERENCES users (userid) ON DELETE CASCADE
				);`,
			//ON DELETE CASCADE: each row in the child table that was associated with the deleted parent row is also deleted
			`CREATE TABLE comments (
				commentid INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
				userid INTEGER NOT NULL, 
				photoid INTEGER NOT NULL,
				commentText TEXT NOT NULL,
				FOREIGN KEY(photoid) REFERENCES photos (photoid),
				FOREIGN KEY(userid) REFERENCES photos (userid)
				);`,
			`CREATE TABLE followers (
				followerid INTEGER NOT NULL,
				followedid INTEGER NOT NULL,
				PRIMARY KEY (followedid, followerid), 
				FOREIGN KEY(followerid) REFERENCES users (userid) ON DELETE CASCADE,
				FOREIGN KEY(followedid) REFERENCES users (userid) ON DELETE CASCADE  
				);`,
			//following table is useless because it will contain the same elements of followers table
			`CREATE TABLE likes (
				photoid INTEGER NOT NULL,
				userid INTEGER NOT NULL,
				PRIMARY KEY (userid, photoid),
				FOREIGN KEY(userid) REFERENCES users (userid) ON DELETE CASCADE,
				FOREIGN KEY(photoid) REFERENCES photos (photoid) ON DELETE CASCADE
				);`,
			`CREATE TABLE banned (
				bannedid INTEGER NOT NULL,
				bannerid INTEGER NOT NULL,
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
