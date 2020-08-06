package handler

import (
	"log"
	"memoapp/model"
	"memoapp/repository"

	"github.com/labstack/echo/v4"
)

func (m *memohandler) CheckCache() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			memo := &model.Memo{}
			err := c.Bind(memo)

			if err != nil {
				c.Logger().Errorf("failed to bind : %v\n", err)
				return err
			}

			if memo.Memo == "Usecache" {
				cache, err := repository.ProvieCache(c)
				if err != nil {
					log.Println("キャッシュインスタンスの取得に失敗しました")
					return err
				}
				ok := cache.Judge()
				if ok {
					m.db = cache
				} else {
					mysql, err := repository.ProvideMysql(c)
					if err != nil {
						log.Println("MySqlインスタンスの取得に失敗しました")
						return err
					}
					m.db = mysql
				}

			} else {
				mysql, err := repository.ProvideMysql(c)
				if err != nil {
					log.Println("MySqlインスタンスの取得に失敗しました")
					return err
				}
				m.db = mysql
			}

			return next(c)
		}
	}
}
