package constants

import (
	"errors"
	"os"
	"strconv"
	"time"

	"webapi/interfaces/graylistx"
)

const (
	cacheProtocol           = "STORE_CACHE_PROTOCOL"
	cacheMaxActive          = "STORE_CACHE_MAX_ACTIVE"
	cacheMaxIdle            = "STORE_CACHE_MAX_IDLE"
	cacheIdleTimeoutSeconds = "STORE_CACHE_IDLE_TIMEOUT_SECONDS"
	cacheHost               = "STORE_CACHE_HOST"
	cachePort               = "STORE_CACHE_PORT"
)

const (
	Production 	= "PRODUCTION"
	Development = "TEST"
	Local				= "UNIT_TESTS"
)

var Env, ErrEnv = getConfig()

func getConfig() (*graylistx.Config, error) {
	host := os.Getenv(cacheHost)
	port := os.Getenv(cachePort)
	protocol := os.Getenv(cacheProtocol)
	maxIdle := os.Getenv(cacheMaxIdle)
	maxActive := os.Getenv(cacheMaxActive)
	idleTimoutSeconds := os.Getenv(cacheIdleTimeoutSeconds)

	if host == "" || port == "" || protocol == "" || maxIdle == "" || maxActive == "" || idleTimoutSeconds == "" {
		return nil, errors.New(
			"cache - getConfigFromEnv - unable to import required evnironment variables",
		)
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

	cache := graylistx.Config{
		Host:        host,
		Port:        portAsInt,
		Protocol:    protocol,
		MaxIdle:     maxIdleAsInt,
		IdleTimeout: idleTimeoutAsTime,
		MaxActive:   maxActiveAsInt,
	}

	return &cache, nil
}