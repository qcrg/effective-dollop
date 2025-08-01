package initiator

import (
	"sync"

	"github.com/qcrg/effective-dollop/config"
)

func get_config() (*config.Config, error) {
	conf, err := config.Load("./config.toml")
	if err != nil {
		return nil, err
	}
	return conf, nil
}

var GetConfig = sync.OnceValues(get_config)
var InitConfig = sync.OnceValue(func() error {
	_, err := GetConfig()
	return err
})
