package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Return the user's photos in reverse chronological order and the user's followers and following
func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")
	userId := extractToken(r.Header.Get("Authorization"))
	requestedUser := ps.ByName("userid")

	isUser, err := rt.db.CheckUser(requestedUser)
	if err != nil {
		//User does not exist
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isUser == false {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	isBanned, err := rt.db.BanCheck(userId, requestedUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isBanned == false {
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
		UserId:    userId,
		Following: following,
		Followers: followers,
		Photos:    photos,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
