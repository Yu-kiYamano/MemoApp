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
	mysql, err := ProvideMysql(c)
	if err != nil {
		c.Logger().Errorf("failed to bind : %v\n", err)
		return nil, err
	}
	cache, err := ProvieCache(c)
	if err != nil {
		c.Logger().Errorf("failed to bind : %v\n", err)
		return nil, err
	}
	c.Logger().Print(cache)
	c.Logger().Print(mysql)
	return cache, nil

}
