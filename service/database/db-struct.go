package database

import "time"

//This file contains all the structures needed for the database package

type User struct {
	UserId    int64
	Username  string
	Followers []int
	Following []int
	Banned    []int
	Photos    []Photo
}

type Comment struct {
	UserId      int64
	PhotoId     int64
	CommentId   int64
	CommentText string
}

type Photo struct {
	PhotoId  int64
	UserId   int64
	Likes    []int
	Comments []Comment
	Date     time.Time
}

type Username struct {
	Username string
}

type UserId struct {
	UserId int64
}

type PhotoId struct {
	PhotoId int64
}

type CommentId struct {
	CommentId int64
}

type CommentText struct {
	CommentText string
}
