package api

import (
	"WASAPhoto/service/api/reqcontext"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
	"path/filepath"
)

const userUploads = "./user_uploads"

// If the user does not exist, it will be created and an identifier is returned.
// If the user exists, the user identifier is returned.
func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, context reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	var userId string
	err := json.NewDecoder(r.Body).Decode(&userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = rt.db.CreateUser(userId)
	if err != nil {
		// User already exists
		err = json.NewEncoder(w).Encode(userId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	if !isValid(userId) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = createFolder(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(User{
		UserId:    userId,
		Following: nil,
		Followers: nil,
		Banned:    nil,
		Photos:    nil,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func isValid(userid string) bool {
	if len(userid) < 3 || len(userid) > 16 {
		return false
	}
	return true
}

// createFolder creates a local folder for the specified userId
func createFolder(userid string) error {
	path := filepath.Join(userUploads, userid)

	// Create the upload directory for the user if it doesn't exist
	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}

	return nil
}
