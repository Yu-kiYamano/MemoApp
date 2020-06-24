package repository

import (
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func SetDB(d *sqlx.DB) { //引数にデータベースとの接続情報を持った構造体を取り,repositoryパッケージのグローバル変数にセット。これでrepositoryパッケージ内でdbアクセスが可能
	db = d
}
