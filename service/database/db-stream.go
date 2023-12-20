package database

// GetStream returns the list of photo of the users followed by an user
func (db *appdbimpl) GetStream(userid string) ([]Photo, error) {
	rows, err := db.c.Query("SELECT * FROM photos WHERE userid IN (SELECT followedid FROM followers WHERE followedid = ?) ORDER BY date DESC", userid)

	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	var stream []Photo

	for rows.Next() {
		var p Photo
		err = rows.Scan(&p.PhotoId, &p.UserId, &p.Date)
		if err != nil {
			return nil, err
		}
		stream = append(stream, p)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return stream, nil
}
