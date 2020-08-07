package handler

import (
	"fmt"
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
			cache, err := repository.ProvieCache(c)
			if err != nil {
				log.Println("キャッシュインスタンスの取得に失敗しました")
				return err
			}
			ok := cache.Judge()
			if ok {
				log.Println("キャッシュにヒットしました")
				m.db = cache
			} else {
				log.Println("キャッシュにヒットしませんでした")
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

func (m *memohandler) SetCache() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err != nil {
				log.Println("エラー")
				return err
			}
			memo := &model.Memo{}
			err = c.Bind(memo)

			if err != nil {
				c.Logger().Errorf("failed to bind : %v\n", err)
				return err
			}
			switch dbType := m.db.(type) {
			case repository.Mysql:
				cache, err := repository.ProvieCache(c)
				if err != nil {
					log.Println("キャッシュインスタンスの取得に失敗しました")
					return err
				}
				m.db = cache

			case repository.Cache:
				break

			default:

				return fmt.Errorf("unexpected type %v ", dbType)
			}

			err = m.db.Set(c, memo)
			if err != nil {
				return err
			}
			log.Println("キャッシュをセットしました")
			return nil
		}
	}
}
