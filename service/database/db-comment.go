package database

// CommentPhoto adds a comment of an user on a photo
func (db *appdbimpl) CommentPhoto(photo PhotoId, user UserId, comment CommentText) (int64, error) {
	id, err := db.c.Exec("INSERT INTO comments (userid, photoid, commentText) VALUES(?, ?, ?) ", user.UserId, photo.PhotoId, comment.CommentText)

	if err != nil {
		return -1, err
	}
	commentid, err := id.LastInsertId()
	return commentid, err
}

// UncommentPhoto removes a comment of an user from a photo
func (db *appdbimpl) UncommentPhoto(photo PhotoId, user UserId, comment CommentId) error {
	_, err := db.c.Exec("DELETE FROM comments WHERE (commentid = ? AND userid = ? AND photoid = ?) ", comment.CommentId, user.UserId, photo.PhotoId)

	if err != nil {
		return err
	}

	return nil
}

// GetPhotoComments returns the list of comments of a photo
func (db *appdbimpl) GetPhotoComments(photo PhotoId) ([]Comment, error) {
	rows, err := db.c.Query("SELECT * FROM comments WHERE photoid = ?", photo.PhotoId)

	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	var comments []Comment

	for rows.Next() {
		var c Comment
		err = rows.Scan(&c.CommentId, &c.CommentText, &c.UserId, &c.PhotoId)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return comments, nil
}
