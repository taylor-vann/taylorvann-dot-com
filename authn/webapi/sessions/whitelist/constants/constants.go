package constants

import (
	"errors"
	"os"
	"strconv"
	"time"

	"webapi/utils/graylistx"
)

const (
	whitelistHost               = "WHITELIST_HOST"
	whitelistPort               = "WHITELIST_PORT"
	whitelistProtocol           = "WHITELIST_PROTOCOL"
	whitelistMaxActive          = "WHITELIST_MAX_ACTIVE"
	whitelistMaxIdle            = "WHITELIST_MAX_IDLE"
	whitelistIdleTimeoutSeconds = "WHITELIST_IDLE_TIMEOUT_SECONDS"
)

const (
	Production  = "PRODUCTION"
	Development = "TEST"
	Local       = "UNIT_TESTS"
)

var Env, ErrEnv = getConfig()

var (
	InvalidConfiguration = errors.New(
		"whitelistx - getConfigFromEnv - unable to import required evnironment variables",
	)
)

func getConfig() (*graylistx.Config, error) {
	host := os.Getenv(whitelistHost)
	port := os.Getenv(whitelistPort)
	protocol := os.Getenv(whitelistProtocol)
	maxIdle := os.Getenv(whitelistMaxIdle)
	maxActive := os.Getenv(whitelistMaxActive)
	idleTimoutSeconds := os.Getenv(whitelistIdleTimeoutSeconds)

	if host == "" || port == "" || protocol == "" || maxIdle == "" || maxActive == "" || idleTimoutSeconds == "" {
		return nil, InvalidConfiguration
	}

	portAsInt, errPort := strconv.Atoi(port)
	if errPort != nil {
		return nil, errPort
	}

	maxIdleAsInt, errIdle := strconv.Atoi(maxIdle)
	if errIdle != nil {
		return nil, errIdle
	}

	maxActiveAsInt, errMaxActive := strconv.Atoi(maxActive)
	if errMaxActive != nil {
		return nil, errMaxActive
	}

	idleTimoutSecondsAsInt, errIdle := strconv.Atoi(idleTimoutSeconds)
	if errIdle != nil {
		return nil, errIdle
	}

	idleTimeoutAsTime := time.Duration(idleTimoutSecondsAsInt) * time.Second

	whitelist := graylistx.Config{
		Host:        host,
		Port:        portAsInt,
		Protocol:    protocol,
		MaxIdle:     maxIdleAsInt,
		IdleTimeout: idleTimeoutAsTime,
		MaxActive:   maxActiveAsInt,
	}

	return &whitelist, nil
}
