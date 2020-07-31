package repository

import (
	"database/sql"
	"log"
	"os"

	"memoapp/model"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Cache struct {
	db *sqlx.DB
}

var _ Database = Cache{}

func ProvieCache(c echo.Context) (Cache, error) {
	dsn := os.Getenv("DSN")            //.envrcのDSNを取得してdsnに代入(dsnとはプログラム側が捜査対象のdbを指定するための識別子)
	db, err := sqlx.Open("mysql", dsn) //("mysql"(ドライバ名),dsn(dsnの名前(26行目で定義))　*sql.DB(つまりdb)を返す )
	if err != nil {
		c.Logger().Errorf("データベース接続に失敗しました。: %v\n", err)
		return Cache{}, err
	}

	if err := db.Ping(); err != nil { //Pingとは対処のコンピュータとネットワークで繋がっているかを確認する時に使うもの
		c.Logger().Errorf("確認できません: %v\n", err)
		return Cache{}, err
	}
	log.Println("データベースに接続しました")

	return Cache{db: db}, nil
}

func (cache Cache) Close() error {
	err := cache.db.Close()
	return err

}

func (cache Cache) Set(c echo.Context, memo *model.Memo) (sql.Result, error) {
	query := `INSERT INTO memos (memo) VALUES (:memo);` //queryにSQL文を代入
	tx, err := cache.db.Beginx()                        //トランザクション開始
	if err != nil {
		c.Logger().Errorf("トランザクションが開始されませんでした: %v\n", err)

		return nil, err
	}
	memo.Memo += "(キャッシュを取得しました。)"
	res, err := tx.NamedExec(query, memo) //queryと構造体を引数に渡してSQLを実行
	if err != nil {
		tx.Rollback()   //エラーが発生した場合はロールバックする
		return nil, err //エラー内容を返す
	}
	tx.Commit()     //成功した場合はコミット
	return res, nil //SQLの実行結果を返す
}

func (cache Cache) Get() ([]*model.Memo, error) {

	return nil, nil
}

func (cache Cache) Delete(c echo.Context, id int) error {
	return nil
}
