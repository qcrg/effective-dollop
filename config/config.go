package config

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

type General struct {
	Domain   string        `toml:"domain"`
	Port     int           `toml:"port"`
	Interval time.Duration `toml:"interval"`
}

type Log struct {
	Level string `toml:"level"`
}

type TLS struct {
	CertPath string `toml:"cert_path"`
	SkeyPath string `toml:"skey_path"`
}

type DB_Postgres struct {
	Name     string `toml:"name"`
	Port     int    `toml:"port"`
	Username string `toml:"username"`
	TlsMode  string `toml:"tls_mode"`
}

type Database struct {
	Postgres DB_Postgres `toml:"postgres"`
}

type Coingecko struct {
	UpdateListInterval time.Duration `toml:"update_list_interval"`
}

type Config struct {
	General   General   `toml:"general"`
	Log       Log       `toml:"log"`
	Tls       TLS       `toml:"tls"`
	Database  Database  `toml:"database"`
	Coingecko Coingecko `toml:"coingecko"`
}

func (*Config) GetEnvConfig() ConfigEnv {
	return ConfigEnv{}
}

const (
	_PG_KEY_PREFIX = "EVDP_DATABASE_POSTGRES_"
	PG_HOST_KEY    = _PG_KEY_PREFIX + "HOST"
	PG_PASSWD_KEY  = _PG_KEY_PREFIX + "PASSWD"
)

type ConfigEnv struct{}

func (ConfigEnv) GetPostgresHost() string {
	return os.Getenv(PG_HOST_KEY)
}

func (ConfigEnv) GetPostgresPassword() string {
	return os.Getenv(PG_PASSWD_KEY)
}

func check_env_value(key string) error {
	val, exists := os.LookupEnv(key)
	if !exists {
		return fmt.Errorf("%s is not defined", key)
	}
	if val == "" {
		return fmt.Errorf("%s is empty", key)
	}
	return nil
}

func check_env() error {
	return errors.Join(
		check_env_value(PG_HOST_KEY),
		check_env_value(PG_PASSWD_KEY),
	)
}

func Load(path string) (*Config, error) {
	conf := Config{}
	_, err := toml.DecodeFile(path, &conf)
	if err != nil {
		return nil, err
	}
	err = check_env()
	if err != nil {
		return nil, err
	}
	return &conf, nil
}
