package api

import (
	"WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Ban an user
func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, context reqcontext.RequestContext) {

	userId := removeBearer(r.Header.Get("Authorization"))

	if isAuth(userId) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	bannedId := ps.ByName("banneduid")

	isBanned, err := rt.db.BanCheck(userId, bannedId)
	if err != nil {
		context.Logger.WithError(err).Error("banUser/BanCheck: error while executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isBanned == true {
		// bannedId is already banned
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = rt.db.BanUser(bannedId, userId)
	if err != nil {
		context.Logger.WithError(err).Error("banUser/BanUser: error while executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
