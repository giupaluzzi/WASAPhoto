package database

//	CreatePhoto creates a photo
func (db *appdbimpl) CreatePhoto(photo Photo) (int, error) {
	_, err := db.c.Exec("INSERT INTO photos (userid, date) VALUES(?, ?) ", photo.UserId, photo.Date)
	
	if err != nil {
		return -1, err
	}

	photoid, err := _.LastInsertId()
	return photoid, err
}

//	DeletePhoto removes a photo
func (db *appdbimpl) DeletePhoto(photo Photo) error {
	_, err := db.c.Exec("DELETE FROM photos WHERE (photoid = ? AND userid = ? AND date = ?) ", photo.PhotoId, photo.UserId, photo.Date)
	
	if err != nil {
		return err
	}

	return nil
}
