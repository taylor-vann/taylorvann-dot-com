// brian taylor vann
// toolbox-go

package pgsqlx

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Config struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	DatabaseName string `json:"database_name"`
}

const (
	connectionString = "postgresql://%s:%s@%s:%d/%s?sslmode=disable"
)

func getConnectionStr(config *Config) string {
	return fmt.Sprintf(
		connectionString,
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DatabaseName,
	)
}

func Create(config *Config) (*pgxpool.Pool, error) {
	if config == nil {
		return nil, errors.New("pgsqlx.Create() - nil config provided")
	}

	connStr := getConnectionStr(config)

	return pgxpool.Connect(context.Background(), connStr)
}
