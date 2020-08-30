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

	"webapi/utils/pgsqlx"
	"webapi/utils/storex/constants"
)

var pgsqlConfig, configErr = getConfigFromEnv()
var pgsqlxInstance, pgsqlxErr = pgsqlx.Create(pgsqlConfig)

func getConfigFromEnv() (*pgsqlx.PGConfig, error) {
	if constants.Host == "" || constants.Port == "" || constants.User == "" || constants.Password == "" || constants.Database == "" {
		return nil, errors.New(
			"storex.getConfigFromEnv() - unable to import required evnironment variables",
		)
	}

	portAsInt, err := strconv.Atoi(constants.Port)
	if err != nil {
		return nil, errors.New(
			"storex.getConfigFromEnv() - could not convert env variable 'port' to int",
		)
	}

	config := pgsqlx.PGConfig{
		Host:         constants.Host,
		Port:         portAsInt,
		Username:     constants.User,
		Password:     constants.Password,
		DatabaseName: constants.Database,
	}

	return &config, nil
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	if pgsqlxErr != nil {
		return nil, errors.New("storex.Exec() - there is not a valid instance of pgsqlx")
	}

	return pgsqlxInstance.DB.Exec(query, args...)
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	if pgsqlxErr != nil {
		return nil, errors.New("storex.Query() - there is not a valid instance of pgsqlx")
	}

	return pgsqlxInstance.DB.Query(query, args...)
}
