package coingecko

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/qcrg/effective-dollop/config"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

type Coingecko struct {
	update_interval time.Duration
	names           map[string]any
	log             zerolog.Logger
}

type Coin struct {
	Id string `json:"id"`
}

func (t *Coingecko) UpdateList() error {
	resp, err := http.DefaultClient.Get("https://api.coingecko.com/api/v3/coins/list")
	if err != nil {
		return fmt.Errorf("Failed to UpdateList: %w", err)
	}
	names := make(map[string]any)
	dec := json.NewDecoder(resp.Body)
	var coins []Coin
	err = dec.Decode(&coins)
	if err != nil {
		return fmt.Errorf("Failed to UpdateList: %w", err)
	}
	for _, coin := range coins {
		names[strings.ToLower(coin.Id)] = nil
	}
	t.names = names
	return nil
}

func (t *Coingecko) UpdateListLoop() {
	ticker := time.NewTicker(t.update_interval)
	defer ticker.Stop()
	for {
		t.log.Debug().Msg("Updating coingecko list...")
		err := t.UpdateList()
		if err != nil {
			t.log.Error().
				Str("tag", "coingecko").
				Err(err).
				Msg("Failed to update list in loop")
		} else {
			t.log.Info().
				Str("tag", "coingecko").
				Msg("Coins list is updated")
		}
		<-ticker.C
	}
}

func (t *Coingecko) Has(name string) bool {
	if len(t.names) == 0 {
		t.log.Error().Msg("Coingecko list is empty. May be not updated?")
	}
	_, ok := t.names[name]
	return ok
}

func NewCoingecko(conf *config.Config, log zerolog.Logger) *Coingecko {
	zlog.Debug().Msg(conf.Coingecko.UpdateListInterval.String())
	res := Coingecko{update_interval: conf.Coingecko.UpdateListInterval, log: log}
	return &res
}
