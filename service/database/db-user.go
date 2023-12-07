package database

//	CreateUser creates an user
func (db *appdbimpl) CreateUser(user User, ) (int, error) {
	_, err := db.c.Exec("INSERT INTO users (username) VALUES(?) ", user.Username)
	
	if err != nil {
		return -1, err
	}

	userid, err := _.LastInsertId()
	return userid, err
}
