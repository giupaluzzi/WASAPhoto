package api

import (
	"WASAPhoto/service/database"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

// Upload a new photo
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	//userId of the logged user
	userId := extractToken(r.Header.Get("Authorization"))
	photoId, err := rt.db.CreatePhoto(database.Photo{
		PhotoId:  0,
		UserId:   userId,
		Likes:    nil,
		Comments: nil,
		Date:     time.Time{},
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(Photo{
		PhotoId:  photoId,
		UserId:   userId,
		Likes:    nil,
		Comments: nil,
		Date:     time.Time{},
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
