package api

import (
	"WASAPhoto/service/api/reqcontext"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// Comment a photo
func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, context reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	userId := removeBearer(r.Header.Get("Authorization"))

	if userId != loggedUser {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	photoId, err := strconv.Atoi(ps.ByName("photoid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	photo, err := rt.db.GetPhoto(photoId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	isBanned, err := rt.db.BanCheck(userId, photo.UserId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isBanned == true {
		w.WriteHeader(http.StatusForbidden)
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
