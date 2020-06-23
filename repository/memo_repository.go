package repository

import (
	"database/sql"

	"memoapp/model"

	"github.com/labstack/echo/v4"
)

func MemoCreate(c echo.Context, memo *model.Memo) (sql.Result, error) {
	query := `INSERT INTO memos (memo) VALUES (:memo);`
	tx, err := db.Beginx()
	if err != nil {
		c.Logger().Errorf("トランザクションが開始されませんでした: %v\n", err) //書き換える

		return nil, err
	}
	res, err := tx.NamedExec(query, memo)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return res, nil
}

func Getmemo() ([]*model.Memo, error) {

	query := `SELECT *FROM memos`
	memos := make([]*model.Memo, 0)
	err := db.Select(&memos, query)

	if err != nil {
		return nil, err
	}
	return memos, nil
}

func MemoDelete(c echo.Context, id int) error {
	query := "DELETE FROM memos WHERE id = ?"
	tx, err := db.Beginx()
	if err != nil {
		c.Logger().Errorf("トランザクションが開始されませんでした: %v\n", err)
		return err
	}
	_, err = tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
