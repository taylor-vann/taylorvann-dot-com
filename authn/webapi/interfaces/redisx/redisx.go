// brian taylor vann
// taylorvann dot com

// Package redisx - utility methods to connect to a redis instance
package redisx

import (
	"errors"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

// RedisConfig - Required properties for redis connection
type RedisConfig struct {
	Host                string
	Port                int
	Protocol            string
	MaxIdle             int
	IdleTimeout         time.Duration
	MaxActive           int
}

// RedisConnection - Reference to our database and config
type RedisConnection struct {
	Store  *redis.Pool
	Config *RedisConfig
}

// Close - RedisConnection associative method, close redis pool connections
func (rdsConn *RedisConnection) Close() (*RedisConnection, error) {
	if rdsConn.Store == nil {
		return rdsConn, errors.New("RedisConnection.Close() - Store is nil")
	}

	err := rdsConn.Store.Close()
	if err != nil {
		return rdsConn, err
	}

	return rdsConn, nil
}

// Create - module function, create a new radis connection through a config
func Create(config *RedisConfig) (*RedisConnection, error) {
	if config == nil {
		return nil, errors.New(
			"redix.Create() - nil config provided",
		)
	}

	portAsStr := strconv.Itoa(config.Port)
	redisAddress := config.Host + ":" + portAsStr

	// attempt to connect to redis through radix
	pool := redis.Pool{
		MaxIdle:     config.MaxIdle,
		IdleTimeout: config.IdleTimeout,
		MaxActive:   config.MaxActive,
		Dial: func() (redis.Conn, error) {
			dial, err := redis.Dial(
				config.Protocol,
				redisAddress,
			)

			if err != nil {
				return nil, err
			}

			return dial, nil
		},
	}

	// return our redis connection interface
	redisConn := RedisConnection{
		Store:  &pool,
		Config: config,
	}

	return &redisConn, nil
}
