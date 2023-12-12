package api

import (
	"WASAPhoto/service/database"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// Return the user's photos in reverse chronological order and the user's followers and following
func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")
	userId, err := strconv.ParseInt(ps.ByName("userid"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	username, err := rt.db.GetUsername(database.UserId{UserId: userId})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	following, err := rt.db.GetFollowing(database.UserId{UserId: userId})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	followers, err := rt.db.GetFollowers(database.UserId{UserId: userId})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	photos, err := rt.db.GetPhotoList(database.UserId{UserId: userId})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Profile{
		UserId:    userId,
		Username:  username,
		Following: following,
		Followers: followers,
		Photos:    photos,
	})
}
