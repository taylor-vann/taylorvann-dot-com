// brian taylor vann
// taylorvann dot com

package whitelistx

import (
	"errors"
	"os"
	"strconv"
	"time"

	"webapi/interfaces/redisx"
)

// Constants for Env Variables
const (
	redisHost               = "HOST_REDIS"
	redisPort               = "PORT_REDIS"
	redisProtocol           = "REDIS_PROTOCOL"
	redisMaxActive          = "REDIS_MAX_ACTIVE"
	redisMaxIdle            = "REDIS_MAX_IDLE"
	redisIdleTimeoutSeconds = "REDIS_IDLE_TIMEOUT_SECONDS"
)

// Redis Constants
const (
	redisGet    = "GET"
	redisSet    = "SET"
	redisExpire = "EXPIRE"
)

// Get config from environemnt variables
func getConfigFromEnv() (*redisx.RedisConfig, error) {
	host := os.Getenv(redisHost)
	port := os.Getenv(redisPort)
	protocol := os.Getenv(redisProtocol)
	maxIdle := os.Getenv(redisMaxIdle)
	maxActive := os.Getenv(redisMaxActive)
	idleTimoutSeconds := os.Getenv(redisIdleTimeoutSeconds)

	if host == "" || port == "" || protocol == "" || maxIdle == "" || maxActive == "" || idleTimoutSeconds == "" {
		return nil, errors.New(
			"redisx - getConfigFromEnv - unable to import required evnironment variables",
		)
	}

	portAsInt, err := strconv.Atoi(port)
	if err != nil {
		return nil, errors.New(
			"redisx - getConfigFromEnv - could not convert env variable 'port' to int",
		)
	}

	maxIdleAsInt, err := strconv.Atoi(maxIdle)
	if err != nil {
		return nil, errors.New(
			"redisx - getConfigFromEnv - could not convert env variable 'numberOfConnections' to int",
		)
	}

	maxActiveAsInt, err := strconv.Atoi(maxActive)
	if err != nil {
		return nil, errors.New(
			"redisx - getConfigFromEnv - could not convert env variable 'numberOfConnections' to int",
		)
	}

	idleTimoutSecondsAsInt, err := strconv.Atoi(idleTimoutSeconds)
	if err != nil {
		return nil, errors.New(
			"redisx - getConfigFromEnv - could not convert env variable 'numberOfConnections' to int",
		)
	}

	idleTimeoutAsTime := time.Duration(idleTimoutSecondsAsInt) * time.Second

	// apply env variables to config
	config := redisx.RedisConfig{
		Host:        host,
		Port:        portAsInt,
		Protocol:    protocol,
		MaxIdle:     maxIdleAsInt,
		IdleTimeout: idleTimeoutAsTime,
		MaxActive:   maxActiveAsInt,
	}

	return &config, nil
}

// get our config
var redisConf, errConfig = getConfigFromEnv()

// create instance of redis pool
var redisInst, errRedisx = redisx.Create(redisConf)

// SetAndExpire - controlled set function
func SetAndExpire(key string, value *[]byte, expireSeconds int) (bool, error) {
	if redisInst == nil {
		return false, errors.New("whitelistx - Set - redisInst is nil")
	}

	return true, nil
}

// Get - controlled get function
func Get(key string) (*[]byte, error) {
	if redisInst == nil {
		return nil, errors.New("whitelistx - Set - redisInst is nil")
	}

	// redisInst.Conn.Do
	return nil, nil
}

// Remove - controlled remove function
func Remove(key string) (bool, error) {
	return false, nil
}
