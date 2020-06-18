package repository

import (
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
