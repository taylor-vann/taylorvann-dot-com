// brian taylor vann
// taylorvann dot com

package storex

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"

	"webapi/interfaces/pgsqlx"
)

// Constants for Env Variables
const (
	pgHost         = "HOST_PGSQL"
	pgPort         = "PORT_PGSQL"
	pgUsername     = "POSTGRES_USER"
	pgPassword     = "POSTGRES_PASSWORD"
	pgDatabaseName = "POSTGRES_DB"
)

func getConfigFromEnv() (*pgsqlx.PGConfig, error) {
	fmt.Println("yo dawwg")
	// get env variables
	host := os.Getenv(pgHost)
	port := os.Getenv(pgPort)
	user := os.Getenv(pgUsername)
	password := os.Getenv(pgPassword)
	database := os.Getenv(pgDatabaseName)

	fmt.Println(host, port, user, password, database)

	if host == "" || port == "" || user == "" || password == "" || database == "" {
		return nil, errors.New(
			"pgsqlx - getConfigFromEnv - unable to import required evnironment variables",
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
	config := pgsqlx.PGConfig{
		Host:         host,
		Port:         portAsInt,
		Username:     user,
		Password:     password,
		DatabaseName: database,
	}

	// return address of config
	return &config, nil
}

var pgsqlConfig, configErr = getConfigFromEnv()
var pgsqlxInstance, pgsqlxErr = pgsqlx.Create(pgsqlConfig)

// Exec - expose Exec method without exposing entire db interface
func Exec(query string, args ...interface{}) (sql.Result, error) {
	if pgsqlxErr != nil {
		return nil, errors.New("storex - Exec - there is not a valid instance of pgsqlx")
	}

	return pgsqlxInstance.DB.Exec(query, args...)
}

// Query - expose a method without exposing entire db interface
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	if pgsqlxErr != nil {
		return nil, errors.New("storex - Query - there is not a valid instance of pgsqlx")
	}

	return pgsqlxInstance.DB.Query(query, args...)
}
