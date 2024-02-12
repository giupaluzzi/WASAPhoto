package database

// GetStream returns the list of photo of the users followed by an user
func (db *appdbimpl) GetStream(userid string) ([]Photo, error) {
	rows, err := db.c.Query("SELECT * FROM photos WHERE userid IN (SELECT followedid FROM followers WHERE followerid = ?) ORDER BY date DESC", userid)

	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	var stream []Photo

	for rows.Next() {
		var p Photo
		err = rows.Scan(&p.PhotoId, &p.UserId, &p.Date, &p.File)
		if err != nil {
			return nil, err
		}

		comments, _ := db.GetPhotoComments(p.PhotoId)
		if err != nil {
			return nil, err
		}

		p.Comments = comments

		likes, _ := db.GetPhotoLikes(p.PhotoId)
		if err != nil {
			return nil, err
		}

		p.Likes = likes

		isBanned, err := db.BanCheck(userid, p.UserId)
		if err != nil {
			return nil, err
		}

		if !isBanned {
			stream = append(stream, p)
		}
	}

	if rows.Err() != nil {
		return nil, err
	}

	return stream, nil
}
