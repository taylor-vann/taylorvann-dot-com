package pgsqlx

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"

	// lib/pq is called for its side-effects against database/sql
	_ "github.com/lib/pq"
)

// PGConfig - Required config to connect to PosgreSQL
type PGConfig struct {
	Host         int    `json:"host"`
	Port         int    `json:"port"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	DatabaseName string `json:"database_name"`
}

// PGConnection - Reference to our database and config
type PGConnection struct {
	db     *sql.DB
	config *PGConfig
}

// Close - PGConnection method, Disconnect from postgres via lib/pq
func (pgconn *PGConnection) Close() (*PGConnection, error) {
	if pgconn.db == nil {
		return nil, nil
	}

	err := pgconn.db.Close()
	if err != nil {
		return pgconn, errors.New(
			"pgsqlx - pgsql_interface - failed to close postgresql connection",
		)
	}

	return pgconn, nil
}

// Constants for Env Variables
const (
	pgHost         = "PG_HOST"
	pgPort         = "PG_PORT"
	pgUsername     = "PG_USERNAME"
	pgPassword     = "PG_PASSWORD"
	pgDatabaseName = "PG_DATABASE"
)

func getConfigFromEnv() (*PGConfig, error) {
	// get env variables
	host := os.Getenv(pgHost)
	port := os.Getenv(pgPort)
	user := os.Getenv(pgUsername)
	password := os.Getenv(pgPassword)
	database := os.Getenv(pgDatabaseName)

	if host == "" || port == "" || user == "" || password == "" || database == "" {
		return nil, errors.New(
			"pgsqlx - getConfigFromEnv - unable to import required evnironment variables",
		)
	}

	// get port string as integer
	hostAsInt, err := strconv.Atoi(host)
	if err != nil {
		return nil, errors.New(
			"pgsqlx - getConfigFromEnv - could not convert env variable 'host' to int",
		)
	}

	// get port string as integer
	portAsInt, err := strconv.Atoi(port)
	if err != nil {
		return nil, errors.New(
			"pgsqlx - getConfigFromEnv - could not convert env variable 'port' to int",
		)
	}

	// apply env variables to config
	config := PGConfig{
		Host:         hostAsInt,
		Port:         portAsInt,
		Username:     user,
		Password:     password,
		DatabaseName: database,
	}

	// return address of config
	return &config, nil
}

// Create - Establish a new connection and return
func Create(config *PGConfig) (*PGConnection, error) {
	db, err := sql.Open("postgres", fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%d port=%d sslmode=disable",
		config.Username, config.Password, config.DatabaseName, config.Host, config.Port)
	)
	if err != nil {
		return nil, err
	}

	pgconn := PGConnection{
		db:     db,
		config: config,
	}

	return &pgconn, nil
}
