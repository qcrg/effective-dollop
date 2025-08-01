package currency

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/qcrg/effective-dollop/api/deps"
	"github.com/qcrg/effective-dollop/api/resps"
)

type remove_req struct {
	Coin string `json:"coin"`
}

func RemoveCoinFromObserve(deps *deps.Deps, c echo.Context) error {
	var req remove_req
	err := c.Bind(&req)
	if err != nil {
		err = c.JSON(
			http.StatusBadRequest,
			resps.MakeGenericErr("Request body must be valid JSON object"),
		)
		if err != nil {
			panic(err)
		}
		return fmt.Errorf("Incorrect request body: %w", err)
	}
	req.Coin = strings.ToLower(req.Coin)

	ok := deps.CG.Has(req.Coin)
	if !ok {
		const reason = "Not found coin with this name"
		err = c.JSON(
			http.StatusBadRequest,
			resps.MakeGenericErr(reason),
		)
		if err != nil {
			panic(err)
		}
		return fmt.Errorf(reason)
	}
	coins := deps.DB.Coins()
	exists, err := coins.ExistsByName(req.Coin)
	if err != nil {
		err = c.NoContent(http.StatusInternalServerError)
		if err != nil {
			panic(err)
		}
		return fmt.Errorf("Failed to check coin existence: %w", err)
	}
	if !exists {
		err := c.NoContent(http.StatusNoContent)
		if err != nil {
			panic(err)
		}
		return nil
	}
	err = coins.SetObserveByName(req.Coin, false)
	if err != nil {
		err = c.NoContent(http.StatusInternalServerError)
		if err != nil {
			panic(err)
		}
		return fmt.Errorf("Failed to set observe: %w", err)
	}
	return c.NoContent(http.StatusNoContent)
}
