package external

import (
	"time"

	"github.com/qcrg/effective-dollop/api/deps"
	"github.com/qcrg/effective-dollop/config"
	"github.com/qcrg/effective-dollop/external/coingecko"
	"github.com/qcrg/effective-dollop/postgres"
	"github.com/rs/zerolog"
)

func UpdateCurrencyLoop(conf *config.Config, deps *deps.Deps) {
	ticker := time.NewTicker(conf.General.Interval)
	defer ticker.Stop()
	for {
		deps.Log.Debug().Msg("Updating currency...")
		coins, err := deps.DB.Coins().GetAllObservable()
		if err != nil {
			deps.Log.Error().Err(err).Msg("Failed to get observable objects")
			<-ticker.C
			continue
		}
		if len(coins) == 0 {
			deps.Log.Warn().Msg("Not found symbols to getting currency")
			<-ticker.C
			continue
		}
		var syms []string
		for _, coin := range coins {
			syms = append(syms, coin.Name)
		}
		currs, err := coingecko.GetCurrencies(syms)
		if err != nil {
			deps.Log.Error().Err(err).Msg("Failed to get currencies from coingecko")
			<-ticker.C
			continue
		}

		have_errors := false
		err_coins := zerolog.Arr()
		for _, coin := range coins {
			_, err = deps.DB.Records().Add(postgres.Record{
				ID:         0,
				CoinID:     coin.ID,
				CurrencyID: 1,
				Timestamp:  time.Now(),
				Value:      int64(currs[coin.Name].Usd * 100),
			})
			if err != nil {
				err_coins.Int(coin.ID)
				have_errors = true
			}
		}
		if have_errors {
			deps.Log.Error().
				Array("failed_coin_ids", err_coins).
				Msg("Failed to update currency")
		}
		deps.Log.Info().Msg("Currency updated")
		<-ticker.C
	}
}
