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
