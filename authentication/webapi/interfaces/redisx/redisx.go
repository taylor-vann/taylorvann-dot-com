package redisx

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

// RedisConfig - Required properties for redis connection
type RedisConfig struct {
	Host                string        `json:"host"`
	Port                int           `json:"port"`
	Protocol            string        `json:"protocol"`
	MaxIdle             int           `json:"max_idle"`
	IdleTimeout         time.Duration `json:"idle_timeout"`
	MaxActive           int           `json:"max_active"`
}

// RedisConnection - Reference to our database and config
type RedisConnection struct {
	Store  redis.Conn
	Config *RedisConfig
}

// Close - RedisConnection associative method, close redis pool connections
func (rdsConn *RedisConnection) Close() (*RedisConnection, error) {
	if rdsConn.Store == nil {
		return rdsConn, errors.New("RedisConnection struct - Close - Store is nil")
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
			"redix - Create - nil config provided",
		)
	}

	portAsStr := strconv.Itoa(config.Port)
	redisAddress := config.Host + ":" + portAsStr

	// attempt to connect to redis through radix
	pool := redis.Pool{
		MaxIdle:     config.MaxIdle,
		IdleTimeout: config.IdleTimeout, // 240 * time.Second,
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

	// redis pool connection
	conn := pool.Get()

	// Test the connection
	result, err := conn.Do("PING")
	if err != nil {
		fmt.Println(result)
		return nil, errors.New("redisx - Create - unable to connect to Redis database")
	}

	// return our redis connection interface
	redisConn := RedisConnection{
		Store:  conn,
		Config: config,
	}

	return &redisConn, nil
}
