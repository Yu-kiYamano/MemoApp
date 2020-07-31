package repository

import (
	"database/sql"
	"memoapp/model"

	"github.com/labstack/echo/v4"
)

type Database interface {
	// Connect() Database
	Set(echo.Context, *model.Memo) (sql.Result, error)
	Get() ([]*model.Memo, error)
	Delete(echo.Context, int) error
	Close() error
}

func ProvideDatabase(c echo.Context) (Database, error) {
	memo := &model.Memo{}
	err := c.Bind(memo)

	if err != nil {
		c.Logger().Errorf("failed to bind : %v\n", err)
		return nil, err
	}

	if memo.Memo == "cache" {
		return ProvieCache(c)
	} else {
		return ProvideMysql(c)
	}
}
