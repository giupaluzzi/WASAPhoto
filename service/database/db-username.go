package database

// SetMyUsername updates logged user's username
func (db *appdbimpl) SetMyUsername(userid string, newUserid string) error {
	_, err := db.c.Exec("UPDATE users SET userid = ? WHERE userid = ?", newUserid, userid)

	if err != nil {
		return err
	}

	return nil
}
