package postgres

import (
	"fmt"

	"github.com/qcrg/effective-dollop/config"
)

func make_connection_string(conf *config.Config) string {
	pg := &conf.Database.Postgres
	ec := conf.GetEnvConfig()
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		pg.Username,
		ec.GetPostgresPassword(),
		ec.GetPostgresHost(),
		pg.Port,
		pg.Name,
		pg.TlsMode,
	)
}
