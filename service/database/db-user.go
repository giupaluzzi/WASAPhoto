package database

// CreateUser creates an user
func (db *appdbimpl) CreateUser(user Username) (int64, error) {
	id, err := db.c.Exec("INSERT INTO users (username) VALUES(?) ", user.Username)

	if err != nil {
		return -1, err
	}

	userid, err := id.LastInsertId()
	return userid, err
}
