package database

import "time"

// This file contains all the structures needed for the database package

type User struct {
	UserId string `json:"userid"`
}

type Comment struct {
	UserId      string `json:"userid"`
	PhotoId     int    `json:"photoid"`
	CommentId   int    `json:"commentid"`
	CommentText string `json:"commentText"`
}

type Photo struct {
	PhotoId  int       `json:"photoid"`
	UserId   string    `json:"userid"`
	Likes    []User    `json:"likes"`
	Comments []Comment `json:"comments"`
	Date     time.Time `json:"date"`
	File     []byte    `json:"file"`
}
