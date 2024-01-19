package api

import (
	"WASAPhoto/service/api/reqcontext"
	"WASAPhoto/service/database"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// Upload a new photo
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, context reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	userId := removeBearer(r.Header.Get("Authorization"))

	if userId != loggedUser {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	creationTime := time.Now()
	photoId, err := rt.db.CreatePhoto(database.Photo{
		UserId:   loggedUser,
		Likes:    nil,
		Comments: nil,
		Date:     creationTime,
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	photoFile, _, err := r.FormFile("photoFile")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create empty file for the data of the photo
	dst, err := os.Create(filepath.Join(userUploads, userId, string(rune(photoId))))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(dst, photoFile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = photoFile.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(Photo{
		PhotoId:  photoId,
		UserId:   userId,
		Likes:    nil,
		Comments: nil,
		Date:     creationTime,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
