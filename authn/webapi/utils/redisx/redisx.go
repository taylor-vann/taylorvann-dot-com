// brian taylor vann
// toolbox-go

package redisx

import (
	"errors"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

type RedisConfig struct {
	Host        string
	Port        int
	Protocol    string
	MaxIdle     int
	IdleTimeout time.Duration
	MaxActive   int
}

type Connection struct {
	Store  *redis.Pool
	Config *RedisConfig
}

func (rdsConn *Connection) Close() (*Connection, error) {
	if rdsConn.Store == nil {
		return rdsConn, errors.New("Connection.Close() - Store is nil")
	}

	err := rdsConn.Store.Close()
	if err != nil {
		return rdsConn, err
	}

	return rdsConn, nil
}

func Create(config *RedisConfig) (*Connection, error) {
	if config == nil {
		return nil, errors.New(
			"redix.Create() - nil config provided",
		)
	}

	portAsStr := strconv.Itoa(config.Port)
	redisAddress := config.Host + ":" + portAsStr

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

	redisConn := Connection{
		Store:  &pool,
		Config: config,
	}

	return &redisConn, nil
}
