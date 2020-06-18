package model

type Memo struct {
	ID    int    `db:"id"`
	Title string `db:"title"`
}
