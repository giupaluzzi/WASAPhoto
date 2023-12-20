package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Unfollow a followed user
func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	userId := extractToken(r.Header.Get("Authorization"))
	userToUnfollow := ps.ByName("followinguid")

	err := rt.db.UnfollowUser(userToUnfollow, userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
