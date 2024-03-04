package dbModel

type Comment struct {
	Text   string `db:"text"`
	User   string `db:"user"`
	PostID int    `db:"postID"`
}
