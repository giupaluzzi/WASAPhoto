package api

import (
	"WASAPhoto/service/api/reqcontext"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Return the user's photos in reverse chronological order and the user's followers and following
func (rt *_router) setUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, context reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	userId := removeBearer(r.Header.Get("Authorization"))

	if userId != loggedUser {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var newUserId UserId
	err := json.NewDecoder(r.Body).Decode(&newUserId)
	if err != nil {
		context.Logger.WithError(err).Error("setUsername/Decode/newUserId: error while decoding json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = rt.db.SetMyUsername(userId, newUserId.UserId)
	if err != nil {
		context.Logger.WithError(err).Error("setUsername/SetMyUsername: error while executing db function")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	loggedUser = newUserId.UserId

	w.WriteHeader(http.StatusNoContent)
}
