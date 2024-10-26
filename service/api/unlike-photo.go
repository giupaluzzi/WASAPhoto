package api

import (
	"WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// Remove a like on a photo
func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, context reqcontext.RequestContext) {

	userId := removeBearer(r.Header.Get("Authorization"))

	if isAuth(userId) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	photoId, err := strconv.Atoi(ps.ByName("photoid"))
	if err != nil {
		context.Logger.WithError(err).Error("unlikePhoto/photoId: error while executing query")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = rt.db.UnlikePhoto(photoId, userId)
	if err != nil {
		context.Logger.WithError(err).Error("unlikePhoto/UnlikePhoto: error while executing db function")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
