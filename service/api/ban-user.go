package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Ban an user
func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")
	userId := extractToken(r.Header.Get("Authorization"))
	bannedId := ps.ByName("banneduid")

	err := rt.db.BanUser(bannedId, userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}