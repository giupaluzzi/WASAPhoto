package database

// CreateUser creates an user
func (db *appdbimpl) CreateUser(userid string) error {
	_, err := db.c.Exec("INSERT INTO users (userid) VALUES(?) ", userid)

	if err != nil {
		return err
	}

	return nil
}

func (db *appdbimpl) CheckUser(userid string) (bool, error) {
	rows, err := db.c.Query("SELECT * FROM users WHERE userid=?", userid)

	if err != nil {
		return false, err
	}
	defer func() {
		_ = rows.Close()
	}()

	if rows.Next() == true {
		return true, nil
	}

	if rows.Err() != nil {
		return false, err
	}

	return false, nil
}
