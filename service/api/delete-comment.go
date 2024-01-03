package api

import (
	"WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// Removes a comment of a photo
func (rt *_router) deleteComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params, context reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	userId := extractToken(r.Header.Get("Authorization"))
	photoId, err := strconv.Atoi(ps.ByName("photoid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	commentId, err := strconv.Atoi(ps.ByName("commentid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = rt.db.UncommentPhoto(photoId, userId, commentId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
