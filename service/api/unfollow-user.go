package api

import (
	"WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Unfollow a followed user
func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, context reqcontext.RequestContext) {

	userId := removeBearer(r.Header.Get("Authorization"))

	if isAuth(userId) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	userToUnfollow := ps.ByName("followinguid")

	err := rt.db.UnfollowUser(userToUnfollow, userId)
	if err != nil {
		context.Logger.WithError(err).Error("unfollowUser/UnfollowUser: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
