package database

// BanUser adds an user to another user's banned list
func (db *appdbimpl) BanUser(newBanned UserId, newBanner UserId) error {
	_, err := db.c.Exec("INSERT INTO banned (banner, banned) VALUES(?, ?) ", newBanner.UserId, newBanned.UserId)

	if err != nil {
		return err
	}

	return nil
}

// UnbanUser removes an user from another user's banned list
func (db *appdbimpl) UnbanUser(oldBanned UserId, oldBanner UserId) error {
	_, err := db.c.Exec("DELETE FROM banned WHERE (bannerid = ? AND bannedid = ?) ", oldBanner.UserId, oldBanned.UserId)

	if err != nil {
		return err
	}

	return nil
}

// BanCheck returns "true" if the user A has been banned by user B, "false" otherwise
// This function is used when an user wants to retrieve information about another user
func (db *appdbimpl) BanCheck(bannedUser UserId, bannerUser UserId) (bool, error) {
	isBanned, err := db.c.Query("SELECT bannedid FROM banned WHERE (bannerid = ? AND bannedid = ?)", bannerUser.UserId, bannedUser.UserId)

	if err != nil {
		return false, err
	}

	if isBanned.Next() == true {
		return true, nil
	}

	return false, nil
}
