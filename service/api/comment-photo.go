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
		context.Logger.WithError(err).Error("commentPhoto/photoId: error while executing query")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	photo, err := rt.db.GetPhoto(photoId)
	if err != nil {
		context.Logger.WithError(err).Error("commentPhoto/GetPhoto: error while executing db function")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	isBanned, err := rt.db.BanCheck(userId, photo.UserId)
	if err != nil {
		context.Logger.WithError(err).Error("commentPhoto/BanCheck: error while executing db function")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isBanned == true {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	commentText := r.URL.Query().Get("commentText")
	if err != nil {
		context.Logger.WithError(err).Error("commentPhoto/commentText: error while executing query")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	commentId, err := rt.db.CommentPhoto(photoId, userId, commentText)
	if err != nil {
		context.Logger.WithError(err).Error("commentPhoto/CommentPhoto: error while executing db function")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(Comment{
		CommentId:   commentId,
		PhotoId:     photoId,
		UserId:      userId,
		CommentText: commentText,
	})

	if err != nil {
		context.Logger.WithError(err).Error("commentPhoto/Encode/Comment: error while encoding json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
