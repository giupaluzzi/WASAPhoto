package database

// FollowUser adds a follower to user
func (db *appdbimpl) FollowUser(newFollowed UserId, newFollower UserId) error {
	_, err := db.c.Exec("INSERT INTO followers (followedid, followerid) VALUES(?, ?) ", newFollowed.UserId, newFollower.UserId)

	if err != nil {
		return err
	}

	return nil
}

// UnfollowUser removes a follower from a user
func (db *appdbimpl) UnfollowUser(oldFollowed UserId, oldFollower UserId) error {
	_, err := db.c.Exec("DELETE FROM followers WHERE (followedid = ? AND followerid = ?) ", oldFollowed.UserId, oldFollower.UserId)

	if err != nil {
		return err
	}

	return nil
}

// GetFollowing returns the list of users followed by an user
func (db *appdbimpl) GetFollowing(user UserId) ([]UserId, error) {
	rows, err := db.c.Query("SELECT followedid FROM followers WHERE followerid = ?", user.UserId)

	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	var following []UserId

	for rows.Next() {
		var p UserId
		err = rows.Scan(&p.UserId)
		if err != nil {
			return nil, err
		}
		following = append(following, p)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return following, nil
}

// Getfollowers returns the list an user's followers
func (db *appdbimpl) GetFollowers(user UserId) ([]UserId, error) {
	rows, err := db.c.Query("SELECT followerid FROM followers WHERE followedid = ?", user.UserId)

	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	var followers []UserId

	for rows.Next() {
		var p UserId
		err = rows.Scan(&p.UserId)
		if err != nil {
			return nil, err
		}
		followers = append(followers, p)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return followers, nil
}
