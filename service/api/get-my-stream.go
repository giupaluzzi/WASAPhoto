package api

import (
	"WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Return the user's photos in reverse chronological order and the user's followers and following
func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, context reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")
	userId := extractToken(r.Header.Get("Authorization"))

	_, err := rt.db.GetStream(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}
