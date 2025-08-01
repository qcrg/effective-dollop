package api

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/qcrg/effective-dollop/api/deps"
	v0 "github.com/qcrg/effective-dollop/api/v0"
)

func Init(deps *deps.Deps, app *echo.Group) error {
	app.GET("/version", get_version)
	app.GET("/health", get_health)

	return errors.Join(
		v0.Init(deps, app),
	)
}
