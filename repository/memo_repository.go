package repository

import (
	"database/sql"

	"memoapp/model"
)

func MemoList() ([]*model.Memo, error) {
	query := `SELECT * FROM memos;`

	var memos []*model.Memo
	if err := db.Select(&memos, query); err != nil {
		return nil, err
	}

	return memos, nil
}

func MemoCreate(memo *model.Memo) (sql.Result, error) {
	query := `INSERT INTO memos(title,content)
	VALUES(:title,:content);`

	tx := db.MustBegin()

	res, err := tx.NamedExec(query, memo)
	if err != nil {
		tx.Rollback()

		return nil, err
	}
	tx.Commit()

	return res, nil

}
