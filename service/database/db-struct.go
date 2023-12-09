package database

import "time"

//This file contains all the structures needed for the database package

type User struct {
	UserId int
	Username string
	Followers []int
	Following []int
	Banned []int
	Photos []Photo
}

type Comment struct{
	UserId int
	PhotoId int
	CommentId int
	CommentText string
}

type Photo struct {
	PhotoId int
	UserId int
	Likes []int
	Comments []Comment
	Date time.Time
}

type Username struct {
	Username string 
}

type UserId struct {
	UserId int
}

type PhotoId struct {
	PhotoId int
}

type CommentId struct {
	CommentId int
}

type CommentText struct{
	CommentText string
}
