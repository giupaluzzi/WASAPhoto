package api

import (
	"WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Unban an user
func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, context reqcontext.RequestContext) {

	userId := removeBearer(r.Header.Get("Authorization"))

	if isAuth(userId) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	bannedId := ps.ByName("banneduid")

	isBanned, err := rt.db.BanCheck(bannedId, userId)
	if err != nil {
		context.Logger.WithError(err).Error("unbanUser/BanCheck: error while executing db function")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isBanned == false {
		// bannedId isn't banned
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = rt.db.UnbanUser(bannedId, userId)
	if err != nil {
		context.Logger.WithError(err).Error("unbanUser/UnbanUser: error while executing db function")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
