package postgres

import (
	"database/sql"
)

type Coins struct {
	db *sql.DB
}

type Coin struct {
	ID      int
	Name    string
	Observe bool
}

func (t Coins) ExistsByName(name string) (bool, error) {
	row := t.db.QueryRow("SELECT COUNT(*) FROM coins WHERE name = $1", name)
	var count int
	err := row.Scan(&count)
	return count > 0, err
}

func (t Coins) Add(name string) (int, error) {
	row := t.db.QueryRow(`
			INSERT INTO coins(name, observe)
			VALUES ($1, $2)
			RETURNING id
		`,
		name,
		true,
	)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func (t Coins) AddOrUpdate(
	name string,
	observe bool,
) (int, error) {
	row := t.db.QueryRow(`
			INSERT INTO coins(name, observe)
			VALUES ($1, $2)
			ON CONFLICT (name)
			DO UPDATE SET observe = $2
			RETURNING id
		`,
		name,
		observe,
	)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (t Coins) FindByName(name string) (*Coin, error) {
	row := t.db.QueryRow(`
			SELECT id, observe
			FROM coins
			WHERE name = $1
		`,
		name,
	)
	coin := Coin{Name: name}
	err := row.Scan(&coin.ID, &coin.Observe)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &coin, err
}

func (t Coins) GetAllObservableSybmols() ([]string, error) {
	rows, err := t.db.Query(`
			SELECT name
			FROM coins
			WHERE observe = TRUE
		`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var names []string
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return nil, err
		}
		names = append(names, name)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return names, nil

}

func (t Coins) GetAllObservable() ([]Coin, error) {
	rows, err := t.db.Query(`
			SELECT id, name
			FROM coins
			WHERE observe = TRUE
		`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var coins []Coin
	for rows.Next() {
		var coin Coin
		err := rows.Scan(&coin.ID, &coin.Name)
		if err != nil {
			return nil, err
		}
		coin.Observe = true
		coins = append(coins, coin)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return coins, nil
}

func (t Coins) SetObserveByName(name string, observe bool) error {
	_, err := t.db.Exec(
		"UPDATE coins SET observe = $1 WHERE name = $2",
		observe,
		name,
	)
	return err
}
