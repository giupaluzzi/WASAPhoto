package api

import (
	"WASAPhoto/service/api/reqcontext"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var loggedUser string

// If the user does not exist, it will be created and an identifier is returned.
// If the user exists, the user identifier is returned.
func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, context reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	var userId UserId
	err := json.NewDecoder(r.Body).Decode(&userId)
	if err != nil {
		context.Logger.WithError(err).Error("doLogin/Decode/UserId: error while decoding json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = rt.db.CreateUser(userId.UserId)
	if err != nil {
		// User already exists
		err = json.NewEncoder(w).Encode(userId)
		if err != nil {
			context.Logger.WithError(err).Error("doLogin/Encode/UserId: error while encoding json")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		loggedUser = userId.UserId
		return
	}

	if !isValid(userId.UserId) {
		context.Logger.WithError(err).Error("doLogin/isValid: userId not valid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	loggedUser = userId.UserId

	err = json.NewEncoder(w).Encode(User{
		UserId:    userId.UserId,
		Following: nil,
		Followers: nil,
		Banned:    nil,
		Photos:    nil,
	})

	if err != nil {
		context.Logger.WithError(err).Error("doLogin/Encode/User: error while encoding json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func isValid(userid string) bool {
	if len(userid) < 3 || len(userid) > 16 {
		return false
	}
	return true
}
