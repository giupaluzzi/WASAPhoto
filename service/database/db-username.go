package database

// SetMyUsername updates logged user's username
func (db *appdbimpl) SetMyUsername(user UserId, newUsername Username) error {
	_, err := db.c.Exec("UPDATE users SET username = ? WHERE userid = ?", newUsername.Username, user.UserId)

	if err != nil {
		return err
	}

	return nil
}
