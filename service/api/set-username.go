package api

import (
	"WASAPhoto/service/database"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// Return the user's photos in reverse chronological order and the user's followers and following
func (rt *_router) setUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")
	userId, err := strconv.ParseInt(ps.ByName("userid"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	username := r.URL.Query().Get("username")
	err = rt.db.SetMyUsername(database.UserId(UserId{UserId: userId}), database.Username{Username: username})
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(username)
}
