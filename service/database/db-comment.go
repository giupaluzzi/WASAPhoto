package database

// CommentPhoto adds a comment of an user on a photo
func (db *appdbimpl) CommentPhoto(photoid int, userid string, commentText string) (int, error) {
	id, err := db.c.Exec("INSERT INTO comments (userid, photoid, commentText) VALUES(?, ?, ?) ", userid, photoid, commentText)

	if err != nil {
		return -1, err
	}
	commentid, err := id.LastInsertId()
	return int(commentid), err
}

// UncommentPhoto removes a comment of an user from a photo
func (db *appdbimpl) UncommentPhoto(photoid int, userid string, commentid int) error {
	_, err := db.c.Exec("DELETE FROM comments WHERE (commentid = ? AND userid = ? AND photoid = ?) ", commentid, userid, photoid)

	if err != nil {
		return err
	}

	return nil
}

// GetPhotoComments returns the list of comments of a photo
func (db *appdbimpl) GetPhotoComments(photoid int) ([]Comment, error) {
	rows, err := db.c.Query("SELECT * FROM comments WHERE photoid = ?", photoid)

	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	var comments []Comment

	for rows.Next() {
		var c Comment
		err = rows.Scan(&c.CommentId, &c.UserId, &c.PhotoId, &c.CommentText)
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
