package api

import (
	"WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Follow a new user
func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, context reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	userId := extractToken(r.Header.Get("Authorization"))
	userToFollow := ps.ByName("followinguid")

	err := rt.db.FollowUser(userToFollow, userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
