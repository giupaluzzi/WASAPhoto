package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Return the user's photos in reverse chronological order and the user's followers and following
func (rt *_router) setUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	userId := extractToken(r.Header.Get("Authorization"))
	var newUserId string
	err := json.NewDecoder(r.Body).Decode(&newUserId)
	err = rt.db.SetMyUsername(userId, newUserId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
}
