package repository

import (
	"log"
	"os"

	"memoapp/model"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Mysql struct {
	db *sqlx.DB
}

var _ Database = Mysql{}

func ProvideMysql(c echo.Context) (Mysql, error) {
	dsn := os.Getenv("DSN")            //.envrcのDSNを取得してdsnに代入(dsnとはプログラム側が捜査対象のdbを指定するための識別子)
	db, err := sqlx.Open("mysql", dsn) //("mysql"(ドライバ名),dsn(dsnの名前(26行目で定義))　*sql.DB(つまりdb)を返す )
	if err != nil {
		c.Logger().Errorf("データベース接続に失敗しました。: %v\n", err)
		return Mysql{}, err
	}

	if err := db.Ping(); err != nil { //Pingとは対処のコンピュータとネットワークで繋がっているかを確認する時に使うもの
		c.Logger().Errorf("確認できません: %v\n", err)
		return Mysql{}, err
	}
	log.Println("MySQLに接続しました")

	return Mysql{db: db}, nil
}

func (m Mysql) Close() error {
	err := m.db.Close()
	return err

}

func (m Mysql) Set(c echo.Context, memo *model.Memo) error {
	query := `INSERT INTO memos (memo) VALUES (:memo);` //queryにSQL文を代入
	tx, err := m.db.Beginx()                            //トランザクション開始
	if err != nil {
		c.Logger().Errorf("トランザクションが開始されませんでした: %v\n", err)

		return err
	}
	res, err := tx.NamedExec(query, memo) //queryと構造体を引数に渡してSQLを実行
	if err != nil {
		tx.Rollback() //エラーが発生した場合はロールバックする
		return err    //エラー内容を返す
	}
	tx.Commit() //成功した場合はコミット

	id, err := res.LastInsertId() //SQL実行結果から作成されたレコードのIDを取得する
	//書き換える(mysqlだけのものにする)
	if err != nil {
		c.Logger().Errorf("failed to get ID : %v\n", err)
		return err
		// return c.JSON(http.StatusInternalServerError,
		// 	MemoAppOutput{Message: "ServerError"}) //→handler
	}
	memo.SetId(int(id))
	return nil //SQLの実行結果を返す

}

func (m Mysql) Get() ([]*model.Memo, error) {

	query := `SELECT *FROM memos`
	memos := make([]*model.Memo, 0)   //クエリ結果を格納する空のスライスを用意
	err := m.db.Select(&memos, query) //クエリ結果を代入するmemoのスライスとクエリを指定してqueryを実行

	if err != nil {
		return nil, err
	}
	return memos, nil
}

func (m Mysql) Delete(c echo.Context, id int) error {
	query := "DELETE FROM memos WHERE id = ?"
	tx, err := m.db.Beginx() //トランザクション開始
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
