package initiator

import (
	"os"
	"sync"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func get_default_logger() (zerolog.Logger, error) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixNano
	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).
		With().
		Str("tag", "evdp").
		Logger()

	conf, err := GetConfig()
	if err != nil {
		panic(err)
	}
	if len(conf.Log.Level) == 0 {
		conf.Log.Level = "info"
	}
	level, err := zerolog.ParseLevel(conf.Log.Level)
	if err != nil {
		return logger, err
	}
	zerolog.SetGlobalLevel(level)
	return logger, nil
}

var GetDefaultLogger = sync.OnceValues(get_default_logger)
var InitLogger = sync.OnceValue(func() (err error) {
	log.Logger, err = GetDefaultLogger()
	return
})
