package api

import (
	"WASAPhoto/service/database"
	"time"
)

type User struct {
	UserId    string           `json:"userid"`
	Following []string         `json:"following"`
	Followers []string         `json:"followers"`
	Banned    []string         `json:"banned"`
	Photos    []database.Photo `json:"photos"`
}

type Profile struct {
	UserId    string           `json:"userid"`
	Following []string         `json:"following"`
	Followers []string         `json:"followers"`
	Photos    []database.Photo `json:"photos"`
}

type Photo struct {
	PhotoId  int                `json:"photoid"`
	UserId   string             `json:"userid"`
	Likes    []string           `json:"likes"`
	Comments []database.Comment `json:"comments"`
	Date     time.Time          `json:"date"`
	File     []byte             `json:"file"`
}

type Comment struct {
	CommentId   int    `json:"commentid"`
	PhotoId     int    `json:"photoid"`
	UserId      string `json:"userid"`
	CommentText string `json:"commentText"`
}
