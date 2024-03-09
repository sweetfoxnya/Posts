package dbModel

type Comment struct {
	Text   string `db:"text"`
	Login  string `db:"login"`
	PostID int    `db:"postID"`
}
