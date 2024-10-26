package api

import (
	"WASAPhoto/service/api/reqcontext"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Return the user's photos in reverse chronological order and the user's followers and following
func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, context reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	userId := removeBearer(r.Header.Get("Authorization"))

	if isAuth(userId) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	stream, err := rt.db.GetStream(userId)
	if err != nil {
		context.Logger.WithError(err).Error("getMyStream/GetStream: error while executing db function")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(stream)
	if err != nil {
		context.Logger.WithError(err).Error("getMyStream/Encode/Stream: error while encoding json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
