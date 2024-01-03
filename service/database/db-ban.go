package database

// BanUser adds an user to another user's banned list
func (db *appdbimpl) BanUser(banned string, banner string) error {
	_, err := db.c.Exec("INSERT INTO banned (banner, banned) VALUES(?, ?) ", banner, banned)

	if err != nil {
		return err
	}

	return nil
}

// UnbanUser removes an user from another user's banned list
func (db *appdbimpl) UnbanUser(exBanned string, exBanner string) error {
	_, err := db.c.Exec("DELETE FROM banned WHERE (bannerid = ? AND bannedid = ?) ", exBanner, exBanned)

	if err != nil {
		return err
	}

	return nil
}

// BanCheck returns "true" if the user A has been banned by user B, "false" otherwise
// This function is used when an user wants to retrieve information about another user
func (db *appdbimpl) BanCheck(banned string, banner string) (bool, error) {
	isBanned, err := db.c.Query("SELECT bannedid FROM banned WHERE (bannerid = ? AND bannedid = ?)", banner, banned)

	if err != nil {
		return false, err
	}

	if isBanned.Next() == true {
		return true, nil
	}

	if isBanned.Err() != nil {
		return false, err
	}

	return false, nil
}
