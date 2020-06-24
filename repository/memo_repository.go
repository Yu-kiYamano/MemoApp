package repository

import (
	"database/sql"

	"memoapp/model"

	"github.com/labstack/echo/v4"
)

func MemoCreate(c echo.Context, memo *model.Memo) (sql.Result, error) {
	query := `INSERT INTO memos (memo) VALUES (:memo);` //queryにSQL文を代入
	tx, err := db.Beginx()                              //トランザクション開始
	if err != nil {
		c.Logger().Errorf("トランザクションが開始されませんでした: %v\n", err)

		return nil, err
	}
	res, err := tx.NamedExec(query, memo) //queryと構造体を引数に渡してSQLを実行
	if err != nil {
		tx.Rollback()   //エラーが発生した場合はロールバックする
		return nil, err //エラー内容を返す
	}
	tx.Commit()     //成功した場合はコミット
	return res, nil //SQLの実行結果を返す
}

func Getmemo() ([]*model.Memo, error) {

	query := `SELECT *FROM memos`
	memos := make([]*model.Memo, 0) //クエリ結果を格納する空のスライスを用意
	err := db.Select(&memos, query) //クエリ結果を代入するmemoのスライスとクエリを指定してqueryを実行

	if err != nil {
		return nil, err
	}
	return memos, nil
}

func MemoDelete(c echo.Context, id int) error {
	query := "DELETE FROM memos WHERE id = ?"
	tx, err := db.Beginx() //トランザクション開始
	if err != nil {
		c.Logger().Errorf("トランザクションが開始されませんでした: %v\n", err)
		return err
	}
	_, err = tx.Exec(query, id) //クエリとパラメータ指定してSQLを実行
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
