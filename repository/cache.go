package repository

import (
	"memoapp/model"

	"github.com/labstack/echo/v4"
)

var LocalCache = map[string]interface{}{
	"メモ一覧": []*model.Memo{
		&model.Memo{
			ID:   1,
			Memo: "cacheMemo",
		},
	},
}

type Cache struct {
}

var _ Database = Cache{}

func ProvieCache(c echo.Context) (Cache, error) {
	// dsn := os.Getenv("DSN")            //.envrcのDSNを取得してdsnに代入(dsnとはプログラム側が捜査対象のdbを指定するための識別子)
	// db, err := sqlx.Open("mysql", dsn) //("mysql"(ドライバ名),dsn(dsnの名前(26行目で定義))　*sql.DB(つまりdb)を返す )
	// if err != nil {
	// 	c.Logger().Errorf("データベース接続に失敗しました。: %v\n", err)
	// 	return Cache{}, err
	// }

	// if err := db.Ping(); err != nil { //Pingとは対処のコンピュータとネットワークで繋がっているかを確認する時に使うもの
	// 	c.Logger().Errorf("確認できません: %v\n", err)
	// 	return Cache{}, err
	// }
	// log.Println("データベースに接続しました")

	return Cache{}, nil
}

// func (cache Cache) Close() error {
// 	err := cache.db.Close()
// 	return err

// }

func (cache Cache) Set(c echo.Context, memo *model.Memo) error {
	LocalCache["メモ一覧"] = memo
	return nil
}

func (cache Cache) Get() ([]*model.Memo, error) {
	GetCache := LocalCache["メモ一覧"]
	return GetCache.([]*model.Memo), nil
}

func (cach Cache) Judge() bool {
	_, ok := LocalCache["メモ一覧"]
	return ok
}

func (cache Cache) Delete(c echo.Context, id int) error {
	return nil
}
