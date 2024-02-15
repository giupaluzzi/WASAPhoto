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

		comments, errC := db.GetPhotoComments(p.PhotoId)
		if errC != nil {
			return nil, errC
		}

		p.Comments = comments

		likes, errL := db.GetPhotoLikes(p.PhotoId)
		if errL != nil {
			return nil, errL
		}

		p.Likes = likes

		// Check if requesting user has been banned by requested user
		loggedIsBanned, errLoggedBan := db.BanCheck(userid, p.UserId)
		if errLoggedBan != nil {
			return nil, errLoggedBan
		}

		if !loggedIsBanned {
			// Check if requested user has been banned by requesting user
			isBanned, errBan := db.BanCheck(p.UserId, userid)
			if errBan != nil {
				return nil, errBan
			}

			if !isBanned {
				stream = append(stream, p)
				
			}
		}

	}

	if rows.Err() != nil {
		return nil, err
	}

	return stream, nil
}
