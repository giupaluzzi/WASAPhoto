package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// Comment a photo
func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	userId := extractToken(r.Header.Get("Authorization"))
	photoId, err := strconv.Atoi(ps.ByName("photoid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	commentText := r.URL.Query().Get("commentText")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	commentId, err := rt.db.CommentPhoto(photoId, userId, commentText)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(Comment{
		CommentId:   commentId,
		PhotoId:     photoId,
		UserId:      userId,
		CommentText: commentText,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}