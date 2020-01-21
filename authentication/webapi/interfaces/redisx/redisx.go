package redisx

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/mediocregopher/radix.v3"
)

// RedisConfig - Required properties for redis connection
type RedisConfig struct {
	Host                int    `json:"host"`
	Port                int    `json:"port"`
	Protocol            string `json:"protocol"`
	NumberOfConnections int    `json:"number_of_connections`
}

// RedisConnection - Reference to our database and config
type RedisConnection struct {
	store  *radix.Pool
	config *RedisConfig
}

// Close - RedisConnection associative method, close redis pool connections
func (rdsConn *RedisConnection) Close() (*RedisConnection, error) {
	err := rdsConn.store.Close()
	if err != nil {
		return rdsConn, err
	}
	return rdsConn, nil
}

// Constants for Env Variables
const (
	redisHost                = "REDIS_HOST"
	redisPort                = "REDIS_PORT"
	redisProtocol            = "REDIS_PROTOCOL"
	redisNumberOfConnections = "REDIS_NUMBER_OF_CONNECTIONS"
)

func getConfigFromEnv() (*RedisConfig, error) {
	// get env variables
	host := os.Getenv(redisHost)
	port := os.Getenv(redisPort)
	protocol := os.Getenv(redisProtocol)
	numberOfConnections := os.Getenv(redisNumberOfConnections)

	if host == "" || port == "" || protocol == "" || numberOfConnections == "" {
		return nil, errors.New(
			"redisx - getConfigFromEnv - unable to import required evnironment variables",
		)
	}

	// get host string as integer
	hostAsInt, err := strconv.Atoi(host)
	if err != nil {
		return nil, errors.New(
			"redisx - getConfigFromEnv - could not convert env variable 'host' to int",
		)
	}

	// get port string as integer
	portAsInt, err := strconv.Atoi(port)
	if err != nil {
		return nil, errors.New(
			"redisx - getConfigFromEnv - could not convert env variable 'port' to int",
		)
	}

	// get numberOfConnections string as integer
	numberOfConnectionsAsInt, err := strconv.Atoi(numberOfConnections)
	if err != nil {
		return nil, errors.New(
			"redisx - getConfigFromEnv - could not convert env variable 'numberOfConnections' to int",
		)
	}

	// apply env variables to config
	config := RedisConfig{
		Host:                hostAsInt,
		Port:                portAsInt,
		Protocol:            protocol,
		NumberOfConnections: numberOfConnectionsAsInt,
	}

	// return address of config
	return &config, nil
}

// Create - module function, create a new radis connection through a config
func Create(config *RedisConfig) (*RedisConnection, error) {
	// concat host, ";", port into a string
	hostAsStr, err := strconv.Itoa(config.Host)
	if err != nil {
		return nil, errors.New(
			"redisx - Create - could not convert env variable 'host' to string",
		)
	}

	portAsStr, err := strconv.Itoa(config.Port)
	if err != nil {
		return nil, errors.New(
			"redisx - Create - could not convert env variable 'port' to string",
		)
	}

	// build our redisAddress
	var strBuilder strings.Builder
	strBuilder.WriteString(hostAsStr)
	strBuilder.WriteString(":")
	strBuilder.WriteString(portAsStr)
	redisAddress := strBuilder.String()

	// attempt to connect to redis through radix
	redisConnectionPool, err := radix.NewPool(
		config.Protocol,
		redisAddress,
		config.NumberOfConnections,
	)
	if err != nil {
		return nil, err
	}

	// return redis connection
	redisConn := RedisConnection{
		store:  redisConnectionPool,
		config: config,
	}

	return &redisConn, nil
}
