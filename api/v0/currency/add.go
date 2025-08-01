package currency

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/qcrg/effective-dollop/api/deps"
	"github.com/qcrg/effective-dollop/api/resps"
)

type add_req struct {
	Coin string `json:"coin"`
}

func AddNewCoinToObserve(deps *deps.Deps, c echo.Context) error {
	var req add_req
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
	_, err = coins.AddOrUpdate(req.Coin, true)
	if err != nil {
		c.NoContent(http.StatusInternalServerError)
		return fmt.Errorf("Failed to AddOrUpdate new coin: %w", err)
	}
	return c.NoContent(http.StatusNoContent)
}
