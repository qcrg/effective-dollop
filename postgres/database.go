package postgres

import (
	"database/sql"

	"github.com/qcrg/effective-dollop/config"
)

type DB struct {
	db *sql.DB

	debug string
}

func (t DB) Coins() Coins {
	return Coins{t.db}
}

func (t DB) Currencies() Currencies {
	return Currencies{t.db}
}

func (t DB) Records() Records {
	return Records{t.db}
}

func (t *DB) Close() error {
	return t.db.Close()
}

func (t *DB) GetDebugStr() string {
	return t.debug
}

func NewDatabase(conf *config.Config) (*DB, error) {
	// FIXME
	if "disable" != conf.Database.Postgres.TlsMode {
		log.Fatal().Msgf("Only 'disable' TLS mod is implemented")
	}

	db, err := sql.Open("postgres", make_connection_string(conf))
	if err != nil {
		return nil, err
	}

	return &DB{db: db, debug: make_connection_string(conf)}, nil
}
