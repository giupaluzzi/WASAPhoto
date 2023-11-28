package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
	"math/rand"
)

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
}

//  Return the user's photos in reverse chronological order and the user's followers and following
func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

}