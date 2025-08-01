package deps

import (
	"github.com/qcrg/effective-dollop/external/coingecko"
	"github.com/qcrg/effective-dollop/postgres"
	"github.com/rs/zerolog"
)

type Deps struct {
	DB  postgres.DB
	Log zerolog.Logger
	CG  *coingecko.Coingecko
}
