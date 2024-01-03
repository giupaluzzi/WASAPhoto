package api

import (
	"WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Unban an user
func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, context reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	userId := extractToken(r.Header.Get("Authorization"))
	bannedId := ps.ByName("banneduid")

	err := rt.db.UnbanUser(bannedId, userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
