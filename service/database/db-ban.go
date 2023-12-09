package database

//	BanUser adds an user to another user's banned list
func (db *appdbimpl) FollowUser(newBanned UserId, newBanner UserId) error {
	_, err := db.c.Exec("INSERT INTO banned (banner, banned) VALUES(?, ?) ", newBanner.UserId, newBanned.UserId)
	
	if err != nil {
		return err
	}

	return nil
}

//	UnbanUser removes an user from another user's banned list
func (db *appdbimpl) UnbanUser(oldBanned UserId, oldBanner UserId) error {
	_, err := db.c.Exec("DELETE FROM banned WHERE (bannerid = ? AND bannedid = ?) ", oldBanner.UserId, oldBanned.UserId)
	
	if err != nil {
		return err
	}

	return nil
}
