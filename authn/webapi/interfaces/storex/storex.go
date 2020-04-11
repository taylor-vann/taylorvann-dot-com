//	brian taylor vann
//	taylorvann dot com

//	Package storex - utility to interface Exec and Query from pgsqlx
//
//	storex provides a seperation between queries and postgres

package storex

import (
	"database/sql"
	"errors"
	"strconv"

	"webapi/interfaces/pgsqlx"
	"webapi/interfaces/storex/constants"
)

// getConfigFromEnv -
func getConfigFromEnv() (*pgsqlx.PGConfig, error) {
	if constants.Host == "" || constants.Port == "" || constants.User == "" || constants.Password == "" || constants.Database == "" {
		return nil, errors.New(
			"storex.getConfigFromEnv() - unable to import required evnironment variables",
		)
	}

	// get port string as integer
	portAsInt, err := strconv.Atoi(constants.Port)
	if err != nil {
		return nil, errors.New(
			"storex.getConfigFromEnv() - could not convert env variable 'port' to int",
		)
	}

	// apply env variables to config
	config := pgsqlx.PGConfig{
		Host:         constants.Host,
		Port:         portAsInt,
		Username:     constants.User,
		Password:     constants.Password,
		DatabaseName: constants.Database,
	}

	// return address of config
	return &config, nil
}

var pgsqlConfig, configErr = getConfigFromEnv()
var pgsqlxInstance, pgsqlxErr = pgsqlx.Create(pgsqlConfig)

// Exec - expose Exec method without exposing entire db interface
func Exec(query string, args ...interface{}) (sql.Result, error) {
	if pgsqlxErr != nil {
		return nil, errors.New("storex.Exec() - there is not a valid instance of pgsqlx")
	}

	return pgsqlxInstance.DB.Exec(query, args...)
}

// Query - expose a method without exposing entire db interface
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	if pgsqlxErr != nil {
		return nil, errors.New("storex.Query() - there is not a valid instance of pgsqlx")
	}

	return pgsqlxInstance.DB.Query(query, args...)
}
