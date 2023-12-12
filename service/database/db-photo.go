package database

// CreatePhoto creates a photo
func (db *appdbimpl) CreatePhoto(photo Photo) (int64, error) {
	id, err := db.c.Exec("INSERT INTO photos (userid, date) VALUES(?, ?) ", photo.UserId, photo.Date)

	if err != nil {
		return -1, err
	}

	photoid, err := id.LastInsertId()
	return photoid, err
}

// DeletePhoto removes a photo
func (db *appdbimpl) DeletePhoto(photo Photo) error {
	_, err := db.c.Exec("DELETE FROM photos WHERE (photoid = ? AND userid = ? AND date = ?) ", photo.PhotoId, photo.UserId, photo.Date)

	if err != nil {
		return err
	}

	return nil
}

// GetPhotoList returns the list of photos of an user
func (db *appdbimpl) GetPhotoList(user UserId) ([]Photo, error) {
	rows, err := db.c.Query("SELECT * FROM photos WHERE userid = ? ORDER BY date DESC", user.UserId)

	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	var photos []Photo

	for rows.Next() {
		var p Photo
		err = rows.Scan(&p.PhotoId, &p.UserId, &p.Date)
		if err != nil {
			return nil, err
		}
		photos = append(photos, p)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return photos, nil
}
