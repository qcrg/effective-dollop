package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func get_health(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
