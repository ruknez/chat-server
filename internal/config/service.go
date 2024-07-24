package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/ruknez/chat-server/internal/db/postgres/entity"
)

const (
	dbName     = "PG_DATABASE_NAME"
	pgUserName = "PG_USER"
	pgPassword = "PG_PASSWORD"
	pgPort     = "PG_PORT"
)

// InitPgConfig вычитывает конфиг посгри.
func InitPgConfig() (*entity.PostgresConfig, error) {
	err := godotenv.Load(configPath)
	if err != nil {
		return nil, err
	}

	conf := entity.PostgresConfig{}
	conf.DBName = os.Getenv(dbName)
	if len(conf.DBName) == 0 {
		return nil, errors.New("pg host not found")
	}

	conf.Port, err = strconv.Atoi(os.Getenv(pgPort))
	if len(conf.DBName) == 0 || err != nil {
		return nil, fmt.Errorf("pg host not found %w", err)
	}

	conf.UserName = os.Getenv(pgUserName)
	if len(conf.DBName) == 0 {
		return nil, errors.New("pg user name not found")
	}

	conf.Password = os.Getenv(pgPassword)
	if len(conf.Password) == 0 {
		return nil, errors.New("pg password name not found")
	}

	conf.Host = "localhost"
	return &conf, nil
}
