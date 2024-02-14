package api

import (
	"WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Follow a new user
func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, context reqcontext.RequestContext) {

	userId := removeBearer(r.Header.Get("Authorization"))

	// if userId != loggedUser {
	if isAuth(userId) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	userToFollow := ps.ByName("followinguid")

	err := rt.db.FollowUser(userToFollow, userId)
	if err != nil {
		context.Logger.WithError(err).Error("followUser/FollowUser: error while executing db function")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
