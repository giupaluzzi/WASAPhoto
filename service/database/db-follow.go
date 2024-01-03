package database

// FollowUser adds a follower to user
func (db *appdbimpl) FollowUser(followed string, follower string) error {
	_, err := db.c.Exec("INSERT INTO followers (followedid, followerid) VALUES(?, ?) ", followed, follower)

	if err != nil {
		return err
	}

	return nil
}

// UnfollowUser removes a follower from a user
func (db *appdbimpl) UnfollowUser(exFollowed string, exFollower string) error {
	_, err := db.c.Exec("DELETE FROM followers WHERE (followedid = ? AND followerid = ?) ", exFollowed, exFollower)

	if err != nil {
		return err
	}

	return nil
}

// GetFollowing returns the list of users followed by an user
func (db *appdbimpl) GetFollowing(userid string) ([]string, error) {
	rows, err := db.c.Query("SELECT followedid FROM followers WHERE followerid = ?", userid)

	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	var following []string

	for rows.Next() {
		var p string
		err = rows.Scan(&p)
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

// GetFollowers returns the list an user's followers
func (db *appdbimpl) GetFollowers(userid string) ([]string, error) {
	rows, err := db.c.Query("SELECT followerid FROM followers WHERE followedid = ?", userid)

	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	var followers []string

	for rows.Next() {
		var p string
		err = rows.Scan(&p)
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
