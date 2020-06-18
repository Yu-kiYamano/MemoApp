package model

type Memo struct {
	ID      int    `db:"id" form:"id"`
	Title   string `db:"title" form:"title"`
	Content string `db:"content" form:"content"`
}
