package api

import (
	"WASAPhoto/service/api/reqcontext"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Return the user's photos in reverse chronological order and the user's followers and following
func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, context reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	userId := removeBearer(r.Header.Get("Authorization"))

	if userId != loggedUser {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	requestedUser := ps.ByName("userid")
	isUser, err := rt.db.CheckUser(requestedUser)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isUser == false {
		// User does not exist
		w.WriteHeader(http.StatusForbidden)
		return
	}

	isBanned, err := rt.db.BanCheck(userId, requestedUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isBanned == true {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	following, err := rt.db.GetFollowing(requestedUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	followers, err := rt.db.GetFollowers(requestedUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	photos, err := rt.db.GetPhotoList(requestedUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(Profile{
		UserId:    requestedUser,
		Following: following,
		Followers: followers,
		Photos:    photos,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
