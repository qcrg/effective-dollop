package initiator

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func init_all() (err error) {
	err = godotenv.Load()
	if err != nil && errors.Is(err, os.ErrNotExist) {
		log.Warn().Str("tag", "init").Msg("File '.env' not found")
		err = nil
	}
	err = InitConfig()
	if err != nil {
		return fmt.Errorf("Failed to initiate Config: %w", err)
	}
	err = InitLogger()
	if err != nil {
		return err
	}
	return err
}

var Init = sync.OnceValue(init_all)
