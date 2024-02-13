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

	// Check if requestedUser exists
	isUser, err := rt.db.CheckUser(requestedUser)
	if err != nil {
		context.Logger.WithError(err).Error("getUserProfile/CheckUser: error executing db function")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isUser == false {
		// User does not exist
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Check if the loggedUser has been banned by the requestedUser
	loggedIsBanned, err := rt.db.BanCheck(userId, requestedUser)
	if err != nil {
		context.Logger.WithError(err).Error("getUserProfile/BanCheck/loggedIsBanned: error while executing db function")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if loggedIsBanned == true {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Check if requestedUser is banned
	isBanned, err := rt.db.BanCheck(requestedUser, userId)
	if err != nil {
		context.Logger.WithError(err).Error("getUserProfile/BanCheck/isBanned: error while executing db function")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isBanned == true {
		w.WriteHeader(http.StatusPartialContent)
		return
	}

	following, err := rt.db.GetFollowing(requestedUser)
	if err != nil {
		context.Logger.WithError(err).Error("getUserProfile/GetFollowing: error while executing db function")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	followers, err := rt.db.GetFollowers(requestedUser)
	if err != nil {
		context.Logger.WithError(err).Error("getUserProfile/GetFollowers: error while executing db function")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	photos, err := rt.db.GetPhotoList(requestedUser)
	if err != nil {
		context.Logger.WithError(err).Error("getUserProfile/GetPhotoList: error while executing db function")
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
		context.Logger.WithError(err).Error("getUserProfile/Encode/Profile: error while encoding json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
