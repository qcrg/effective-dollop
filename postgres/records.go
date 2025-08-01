package postgres

import (
	"database/sql"
	"time"
)

type Records struct {
	db *sql.DB
}

type Record struct {
	ID         int64
	CoinID     int
	CurrencyID int
	Timestamp  time.Time
	Value      int64
}

// ignores ID from record
func (t Records) Add(record Record) (int, error) {
	row := t.db.QueryRow(
		"INSERT INTO records(coin_id, currency_id, timestamp, value) VALUES ($1, $2, $3, $4) RETURNING id",
		record.CoinID,
		record.CurrencyID,
		record.Timestamp,
		record.Value,
	)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (t Records) FindValueFromNearestTimestamp(
	coin_id int,
	timestamp time.Time,
) (*int64, *time.Time, error) {
	row := t.db.QueryRow(`
			SELECT value, timestamp
			FROM records
			WHERE coin_id = $1
			ORDER BY ABS(EXTRACT(EPOCH FROM (timestamp - $2)))
			LIMIT 1
		`,
		coin_id,
		timestamp,
	)
	var value int64
	var tstamp time.Time
	err := row.Scan(&value, &tstamp)
	if err == sql.ErrNoRows {
		return nil, nil, nil
	}
	return &value, &tstamp, err
}
