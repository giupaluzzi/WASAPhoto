package database

//	BanUser adds an user to another user's banned list
func (db *appdbimpl) FollowUser(newBanned User, newBanner User) error {
	_, err := db.c.Exec("INSERT INTO banned (banner, banned) VALUES(?, ?) ", newBanner.UserId, newBanned.UserId)
	
	if err != nil {
		return err
	}

	return nil
}

//	UnbanUser removes an user from another user's banned list
func (db *appdbimpl) UnbanUser(oldBanned User, oldBanner User) error {
	_, err := db.c.Exec("DELETE FROM banned WHERE (bannerid = ? AND bannedid = ?) ", oldBanner.UserId, oldBanned.UserId)
	
	if err != nil {
		return err
	}

	return nil
}
