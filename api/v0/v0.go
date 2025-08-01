package v0

import (
	"github.com/labstack/echo/v4"
	"github.com/qcrg/effective-dollop/api/deps"
	"github.com/qcrg/effective-dollop/api/v0/currency"
)

const Version = "0.0.8"

type handler = func(*deps.Deps, echo.Context) error

func bind(deps *deps.Deps, fn handler) echo.HandlerFunc {
	return func(c echo.Context) error {
		return fn(deps, c)
	}
}

func Init(deps *deps.Deps, app *echo.Group) error {
	crg := app.Group("/currency")

	crg.POST("/add", bind(deps, currency.AddNewCoinToObserve))
	crg.POST("/remove", bind(deps, currency.RemoveCoinFromObserve))
	crg.GET("/price", bind(deps, currency.GetPrice))

	return nil
}
