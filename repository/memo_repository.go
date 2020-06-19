package repository

import (
	"database/sql"
	"math"

	"memoapp/model"
)

func MemoCreate(memo *model.Memo) (sql.Result, error) {
	// クエリ文字列を生成します。
	query := `INSERT INTO memos (memo)
	VALUES (:memo);`

	// トランザクションを開始します。
	tx := db.MustBegin()

	// クエリ文字列と構造体を引数に渡して SQL を実行します。
	// クエリ文字列内の「:title」「:body」「:created」「:updated」は構造体の値で置換されます。
	// 構造体タグで指定してあるフィールドが対象となります。（`db:"title"` など）
	res, err := tx.NamedExec(query, memo)
	if err != nil {
		// エラーが発生した場合はロールバックします。
		tx.Rollback()

		// エラー内容を返却します。
		return nil, err
	}

	// SQL の実行に成功した場合はコミットします。
	tx.Commit()

	// SQL の実行結果を返却します。
	return res, nil
}

func MemoListByCursor(cursor int) ([]*model.Memo, error) {
	// 引数で渡されたカーソルの値が 0 以下の場合は、代わりに int 型の最大値で置き換えます。
	if cursor <= 0 {
		cursor = math.MaxInt32
	}

	// ID の降順に記事データを 10 件取得するクエリ文字列を生成します。
	query := `SELECT *
	FROM memos
	WHERE id < ?
	ORDER BY id desc
	LIMIT 10`

	// クエリ結果を格納するスライスを初期化します。
	// 10 件取得すると決まっているため、サイズとキャパシティを指定しています。
	memos := make([]*model.Memo, 0, 10)

	// クエリ結果を格納する変数、クエリ文字列、パラメータを指定してクエリを実行します。
	if err := db.Select(&memos, query, cursor); err != nil {
		return nil, err
	}

	return memos, nil
}

func MemoDelete(id int) error {
	query := "DELETE FROM memos WHERE id = ?"

	tx := db.MustBegin()

	if _, err := tx.Exec(query, id); err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit()
}
