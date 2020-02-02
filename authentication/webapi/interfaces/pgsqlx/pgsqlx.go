package pgsqlx

import (
	"database/sql"
	"errors"
	"fmt"
	// lib/pq is called for its side-effects against database/sql
	_ "github.com/lib/pq"
)

// PGConfig - Required config to connect to PosgreSQL
type PGConfig struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	DatabaseName string `json:"database_name"`
}

// PGConnection - Reference to our database and config
type PGConnection struct {
	DB     *sql.DB
	Config *PGConfig
}

// Close - PGConnection method, Disconnect from postgres via lib/pq
func (pgconn *PGConnection) Close() (*PGConnection, error) {
	if pgconn.DB == nil {
		return nil, nil
	}

	err := pgconn.DB.Close()
	if err != nil {
		return pgconn, errors.New(
			"pgsqlx - pgsql_interface - failed to close postgresql connection",
		)
	}

	return pgconn, nil
}

// Create - Establish a new connection and return
func Create(config *PGConfig) (*PGConnection, error) {
	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=disable",
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
