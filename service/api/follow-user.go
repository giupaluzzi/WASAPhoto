package api

import (
	"WASAPhoto/service/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// Follow a new user
func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")
	userId, err := strconv.ParseInt(ps.ByName("userid"), 10, 64)

	var userToFollow int64
	//	userToFollow is the logged user
	//

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = rt.db.FollowUser(database.UserId{UserId: userToFollow}, database.UserId{UserId: userId})
	w.WriteHeader(http.StatusNoContent)
}
