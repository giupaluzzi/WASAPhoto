package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Upload a new photo
func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	var userId string
	err := json.NewDecoder(r.Body).Decode(&userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = rt.db.CreateUser(userId)
	if err != nil {
		//user already exists
		err = json.NewEncoder(w).Encode(userId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	if !isValid(userId) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(User{
		UserId:    userId,
		Following: nil,
		Followers: nil,
		Banned:    nil,
		Photos:    nil,
	})

	if err != nil {
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
