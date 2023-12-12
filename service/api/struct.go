package api

import (
	"WASAPhoto/service/database"
	"time"
)

type User struct {
	UserId    int64             `json:"userid"`
	Username  string            `json:"username"`
	Following []database.UserId `json:"following"`
	Followers []database.UserId `json:"followers"`
	Banned    []database.UserId `json:"banned"`
	Photos    []database.Photo  `json:"photos"`
}

type UserId struct {
	UserId int64 `json:"userid"`
}

type Profile struct {
	UserId    int64             `json:"userid"`
	Username  string            `json:"username"`
	Following []database.UserId `json:"following"`
	Followers []database.UserId `json:"followers"`
	Photos    []database.Photo  `json:"photos"`
}

type Username struct {
	Username string `json:"username"`
}

type Photo struct {
	PhotoId  int64              `json:"photoid"`
	UserId   int64              `json:"userid"`
	Likes    []database.UserId  `json:"likes"`
	Comments []database.Comment `json:"comments"`
	Date     time.Time          `json:"date"`
}

type PhotoId struct {
	PhotoId int64 `json:"photoid"`
}

type Comment struct {
	CommentId   int64  `json:"commentid"`
	PhotoId     int64  `json:"photoid"`
	UserId      int64  `json:"userid"`
	CommentText string `json:"commentText"`
}

type CommentId struct {
	CommentId int64 `json:"commentid"`
}
