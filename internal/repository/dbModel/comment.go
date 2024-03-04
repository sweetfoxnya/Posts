package dbModel

type Comment struct {
	Text   string `db:"text"`
	UserID int    `db:"userID"`
	PostID int    `db:"postID"`
}
