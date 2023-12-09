package database

//	GetStream returns the list of photo of the users followed by an user
func (db *appdbimpl) GetStream(user UserId) ([]Photo, error){
	rows, err := db.c.Query("SELECT * FROM photos WHERE userid IN (SELECT followedid FROM followers WHERE followedid = ?) ORDER BY date DESC", user.UserId)
	
	if err != nil {
		return nil, err
	}

	defer func {
		_ = rows.Close()
	}()

	var stream []Photo
	
	for r := range(rows) {
		var p Photo
		err = r.Scan(&p.PhotoId, &p.UserId, &p.Date)
		if err != nil {
			return nil, err
		}
		stream.append(stream, p)
	}

	return stream, nil
}
