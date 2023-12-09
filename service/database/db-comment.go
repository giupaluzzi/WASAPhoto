package database

//	CommentPhoto adds a comment of an user on a photo
func (db *appdbimpl) CommentPhoto(photo PhotoId, user UserId, comment CommentText) (int, error) {
	_, err := db.c.Exec("INSERT INTO comments (userid, photoid, commentText) VALUES(?, ?, ?) ", user.UserId, photo.PhotoId, comment.CommentText)
	
	if err != nil {
		return -1, err
	}
	commentid, err := _.LastInsertId()
	return commentid, err
}

//	UncommentPhoto removes a comment of an user from a photo
func (db *appdbimpl) UncommentPhoto(photo PhotoId, user UserId, comment CommentId) error {
	_, err := db.c.Exec("DELETE FROM comments WHERE (commentid = ? AND userid = ? AND photoid = ?) ", comment.CommentId, user.UserId, photo.PhotoId)
	
	if err != nil {
		return err
	}

	return nil
}
