package api

import (
	"WASAPhoto/service/api/reqcontext"
	"WASAPhoto/service/database"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Upload a new photo
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, context reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	userId := removeBearer(r.Header.Get("Authorization"))

	if isAuth(userId) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	file, err := ioutil.ReadAll(r.Body)
	if err != nil {
		context.Logger.WithError(err).Error("uploadPhoto/ReadAll: error while receiving binary file")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	contentType := r.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "image/") {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	creationTime := time.Now()
	photoId, err := rt.db.CreatePhoto(database.Photo{
		UserId:   userId,
		Likes:    nil,
		Comments: nil,
		Date:     creationTime,
		File:     file,
	})
	if err != nil {
		context.Logger.WithError(err).Error("uploadPhoto/CreatePhoto: error while executing db function")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(Photo{
		PhotoId:  photoId,
		UserId:   userId,
		Likes:    nil,
		Comments: nil,
		Date:     creationTime,
		File:     file,
	})

	if err != nil {
		context.Logger.WithError(err).Error("uploadPhoto/Encode/Photo: error while encoding json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
