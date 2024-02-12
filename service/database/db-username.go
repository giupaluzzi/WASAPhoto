package database

import "errors"

// SetMyUsername updates logged user's username
func (db *appdbimpl) SetMyUsername(userid string, newUserid string) error {

	// Check if the newUserid is already taken
	var count int
	err := db.c.QueryRow("SELECT COUNT(*) FROM users WHERE userid = ?", newUserid).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("newUserid already in use")
	}

	_, err = db.c.Exec("PRAGMA foreign_keys = OFF")

	_, err = db.c.Exec("UPDATE users SET userid = ? WHERE userid = ?", newUserid, userid)
	if err != nil {
		return err
	}

	_, err = db.c.Exec("UPDATE photos SET userid = ? WHERE userid = ?", newUserid, userid)
	if err != nil {
		return err
	}

	_, err = db.c.Exec("UPDATE comments SET userid = ? WHERE userid = ?", newUserid, userid)
	if err != nil {
		return err
	}

	_, err = db.c.Exec("UPDATE followers SET followerid = ? WHERE followerid = ?", newUserid, userid)
	if err != nil {
		return err
	}

	_, err = db.c.Exec("UPDATE followers SET followedid = ? WHERE followedid = ?", newUserid, userid)
	if err != nil {
		return err
	}

	_, err = db.c.Exec("UPDATE likes SET userid = ? WHERE userid = ?", newUserid, userid)
	if err != nil {
		return err
	}

	_, err = db.c.Exec("UPDATE banned SET bannerid = ? WHERE bannerid = ?", newUserid, userid)
	if err != nil {
		return err
	}

	_, err = db.c.Exec("UPDATE banned SET bannedid = ? WHERE bannedid = ?", newUserid, userid)
	if err != nil {
		return err
	}

	_, err = db.c.Exec("PRAGMA foreign_keys = ON")

	return nil
}
