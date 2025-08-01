package currency

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/qcrg/effective-dollop/api/deps"
	"github.com/qcrg/effective-dollop/api/resps"
)

type price_req struct {
	Coin      string `json:"coin"`
	Timestamp int64  `json:"timestamp"`
}

func GetPrice(deps *deps.Deps, c echo.Context) error {
	var err error
	var req price_req
	req.Coin = strings.ToLower(c.FormValue("coin"))
	req.Timestamp, err = strconv.ParseInt(c.FormValue("timestamp"), 10, 64)
	if err != nil {
		err = c.JSON(
			http.StatusBadRequest,
			resps.MakeGenericErr("Timestamp must be a number"),
		)
		if err != nil {
			panic(err)
		}
		return fmt.Errorf("Incorrect timestamp: %w", err)
	}

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

	coin, err := deps.DB.Coins().FindByName(req.Coin)
	if err != nil {
		c.NoContent(http.StatusInternalServerError)
		return fmt.Errorf("Failed to get coin by name from db: %w", err)
	}
	if coin == nil {
		err = c.JSON(
			http.StatusNotFound,
			resps.MakeGenericErr("Unsaved prices"),
		)
		if err != nil {
			panic(err)
		}
		return fmt.Errorf("Coin not saved with this name: %s", req.Coin)
	}

	price, price_tstamp, err := deps.DB.Records().FindValueFromNearestTimestamp(
		coin.ID,
		time.Unix(req.Timestamp, 0),
	)
	if err != nil {
		c.NoContent(http.StatusInternalServerError)
		return fmt.Errorf("Failed to find nearest price from db: %w", err)
	}

	if price == nil {
		err = c.JSON(
			http.StatusNotFound,
			resps.MakeGenericErr("Unsaved prices"),
		)
		if err != nil {
			panic(err)
		}
		return fmt.Errorf("Price for coin '%s' not found", req.Coin)
	}

	err = c.JSON(http.StatusOK, echo.Map{
		"coin":      req.Coin,
		"timestamp": price_tstamp.Unix(),
		"price":     float64(*price) / 100,
	})
	if err != nil {
		panic(err)
	}

	return nil
}
