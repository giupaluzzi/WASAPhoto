package database

// CreatePhoto creates a photo
func (db *appdbimpl) CreatePhoto(photo Photo) (int, error) {
	id, err := db.c.Exec("INSERT INTO photos (userid, date) VALUES(?, ?) ", photo.UserId, photo.Date)

	if err != nil {
		return -1, err
	}

	photoid, err := id.LastInsertId()
	return int(photoid), err
}

// DeletePhoto removes a photo
func (db *appdbimpl) DeletePhoto(photoid int, userid string) error {
	_, err := db.c.Exec("DELETE FROM photos WHERE (photoid = ? AND userid = ?) ", photoid, userid)

	if err != nil {
		return err
	}

	return nil
}

// GetPhotoList returns the list of photos of an user
func (db *appdbimpl) GetPhotoList(userid string) ([]Photo, error) {
	rows, err := db.c.Query("SELECT * FROM photos WHERE userid = ? ORDER BY date DESC", userid)

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
