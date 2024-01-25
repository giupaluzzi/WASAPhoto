package api

import (
	"WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// Removes a comment of a photo
func (rt *_router) deleteComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params, context reqcontext.RequestContext) {

	userId := removeBearer(r.Header.Get("Authorization"))

	if userId != loggedUser {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	photoId, err := strconv.Atoi(ps.ByName("photoid"))
	if err != nil {
		context.Logger.WithError(err).Error("deleteComment/photoId: error while executing query")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	commentId, err := strconv.Atoi(ps.ByName("commentid"))
	if err != nil {
		context.Logger.WithError(err).Error("deleteComment/commentId: error while executing query")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	authorCheck, err := rt.db.GetPhoto(photoId)
	if authorCheck.UserId != userId {
		context.Logger.WithError(err).Error("deleteComment/GetPhoto: error while executing db function")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = rt.db.UncommentPhoto(photoId, userId, commentId)
	if err != nil {
		context.Logger.WithError(err).Error("deleteComment/UncommentPhoto: error while executing db function")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
