package database

//	FollowUser adds a follower to user
func (db *appdbimpl) FollowUser(newFollowed User, newFollower User) error {
	_, err := db.c.Exec("INSERT INTO followers (followedid, followerid) VALUES(?, ?) ", newFollowed.UserId, newFollower.UserId)
	
	if err != nil {
		return err
	}

	return nil
}

//	UnfollowUser removes a follower from a user
func (db *appdbimpl) UnfollowUser(oldFollowed User, oldFollower User) error {
	_, err := db.c.Exec("DELETE FROM followers WHERE (followedid = ? AND followerid = ?) ", oldFollowed.UserId, oldFollower.UserId)
	
	if err != nil {
		return err
	}

	return nil
}
