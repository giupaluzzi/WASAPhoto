package database

import "time"

//This file contains all the structures needed for the database package

type User struct {
	UserId    string  `json:"userid"`
	Followers []int   `json:"followers"`
	Following []int   `json:"following"`
	Banned    []int   `json:"banned"`
	Photos    []Photo `json:"photos"`
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
	Likes    []int     `json:"likes"`
	Comments []Comment `json:"comments"`
	Date     time.Time `json:"date"`
}
