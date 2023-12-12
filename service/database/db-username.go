package database

// SetMyUsername updates logged user's username
func (db *appdbimpl) SetMyUsername(user UserId, newUsername Username) error {
	_, err := db.c.Exec("UPDATE users SET username = ? WHERE userid = ?", newUsername.Username, user.UserId)

	if err != nil {
		return err
	}

	return nil
}

// GetUsername returns an user's username
func (db *appdbimpl) GetUsername(user UserId) (string, error) {
	var username string

	err := db.c.QueryRow("SELECT username FROM users WHERE userid = ?", user.UserId).Scan(&username)

	if err != nil {
		return "", err
	}

	return username, nil
}
