package api

import (
	"WASAPhoto/service/api/reqcontext"
	"WASAPhoto/service/database"
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Return the user's photos in reverse chronological order and the user's followers and following
func (rt *_router) setUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, context reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	userId := removeBearer(r.Header.Get("Authorization"))

	if isAuth(userId) {
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
		if errors.Is(err, database.ErrorUseridExists) {
			w.WriteHeader(http.StatusConflict)
			return
		} else {
			context.Logger.WithError(err).Error("setUsername/SetMyUsername: error while executing db function")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
