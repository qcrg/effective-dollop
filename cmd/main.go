package main

import (
	"fmt"
	glog "log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/qcrg/effective-dollop/api"
	"github.com/qcrg/effective-dollop/api/deps"
	"github.com/qcrg/effective-dollop/config"
	"github.com/qcrg/effective-dollop/external"
	"github.com/qcrg/effective-dollop/external/coingecko"
	"github.com/qcrg/effective-dollop/postgres"
	"github.com/qcrg/effective-dollop/utils/initiator"
	"github.com/rs/zerolog/log"
)

func debug() {
	conf, _ := initiator.GetConfig()
	db, err := postgres.NewDatabase(conf)
	if err != nil {
		log.Fatal().Err(err).Msg("db")
	}
	coins, err := db.Coins().GetAllObservable()
	if err != nil {
		log.Fatal().Err(err).Msg("coins")
	}
	for _, coin := range coins {
		fmt.Printf("{id: %d, name: %s}\n", coin.ID, coin.Name)
	}

	os.Exit(0)
}

func make_deps(conf *config.Config) deps.Deps {
	db, err := postgres.NewDatabase(conf)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to make deps")
	}
	log, _ := initiator.GetDefaultLogger()
	return deps.Deps{
		DB:  *db,
		Log: log,
		CG:  coingecko.NewCoingecko(conf, log),
	}
}

func main() {
	var err error

	err = initiator.Init()
	if err != nil {
		glog.Fatal(fmt.Errorf("Failed to initiate system: %w", err))
		return
	}
	conf, _ := initiator.GetConfig()

	deps := make_deps(conf)

	// debug()

	app := echo.New()
	app.HideBanner = true
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	group := app.Group("")
	if group == nil {
		log.Fatal().Msg("Failed to create main group")
	}
	err = api.Init(&deps, group)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to init API")
	}

	go deps.CG.UpdateListLoop()
	go external.UpdateCurrencyLoop(conf, &deps)

	err = app.StartTLS(
		fmt.Sprintf("%s:%d", conf.General.Domain, conf.General.Port),
		conf.Tls.CertPath,
		conf.Tls.SkeyPath,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start http server")
	}
	log.Info().Msg("Stopping server")
}
