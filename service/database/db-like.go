package database

//	LikePhoto adds a like of an user on a photo
func (db *appdbimpl) LikePhoto(photo Photo, user User) error {
	_, err := db.c.Exec("INSERT INTO likes (userid, photoid) VALUES(?, ?) ",  user.UserId, photo.PhotoId)
	
	if err != nil {
		return err
	}

	return nil
}

//	UnlikePhoto removes a like of an user from a photo
func (db *appdbimpl) UnlikePhoto(photo Photo, user User) error {
	_, err := db.c.Exec("DELETE FROM likes WHERE (userid = ? AND photoid = ?) ", user.UserId, photo.PhotoId)
	
	if err != nil {
		return err
	}

	return nil
}
