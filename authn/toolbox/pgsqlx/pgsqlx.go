// brian taylor vann
// toolbox-go

package pgsqlx

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
	"toolbox/pgsqlx/constants"
)

type PGConfig struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	DatabaseName string `json:"database_name"`
}

type PGConnection struct {
	DB     *sql.DB
	Config *PGConfig
}

func (pgconn *PGConnection) Close() (*PGConnection, error) {
	if pgconn.DB == nil {
		return nil, nil
	}

	err := pgconn.DB.Close()
	if err != nil {
		return pgconn, errors.New(
			"pgsqlx.PGConnection.Close() - failed to close postgresql connection",
		)
	}

	return pgconn, nil
}

func Create(config *PGConfig) (*PGConnection, error) {
	if config == nil {
		return nil, errors.New("pgsqlx.Create() - nil config provided")
	}

	connStr := fmt.Sprintf(
		constants.ConnectionString,
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DatabaseName,
	)

	db, err := sql.Open(
		"postgres",
		connStr,
	)

	if err != nil {
		return nil, err
	}

	pgconn := PGConnection{
		DB:     db,
		Config: config,
	}

	return &pgconn, nil
}
