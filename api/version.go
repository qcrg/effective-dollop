package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	v0 "github.com/qcrg/effective-dollop/api/v0"
)

func get_version(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"version": v0.Version})
}
