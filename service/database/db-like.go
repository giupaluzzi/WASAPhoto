package database

// LikePhoto adds a like of an user on a photo
func (db *appdbimpl) LikePhoto(photoid int, userid string) error {
	_, err := db.c.Exec("INSERT INTO likes (userid, photoid) VALUES(?, ?) ", userid, photoid)

	if err != nil {
		return err
	}

	return nil
}

// UnlikePhoto removes a like of an user from a photo
func (db *appdbimpl) UnlikePhoto(photoid int, userid string) error {
	_, err := db.c.Exec("DELETE FROM likes WHERE (userid = ? AND photoid = ?) ", userid, photoid)

	if err != nil {
		return err
	}

	return nil
}

// GetPhotoLikes returns the list of userId who liked the photo
func (db *appdbimpl) GetPhotoLikes(photoid int) ([]User, error) {
	rows, err := db.c.Query("SELECT userid FROM likes WHERE photoid = ?", photoid)

	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	var likes []User

	for rows.Next() {
		var l User
		err = rows.Scan(&l.UserId)
		if err != nil {
			return nil, err
		}
		likes = append(likes, l)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return likes, nil
}
