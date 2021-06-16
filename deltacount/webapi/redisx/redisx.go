// brian taylor vann
// redisx

package redisx

import (
	"errors"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

type Config struct {
	Host        string
	Port        int
	Protocol    string
	MaxIdle     int
	IdleTimeout time.Duration
	MaxActive   int
}

const (
	DELIMITER = ":"
)

func Create(config *Config) (*redis.Pool, error) {
	if config == nil {
		return nil, errors.New(
			"redix.Create() - nil config provided",
		)
	}

	redisAddress := fmt.Sprint(config.Host, DELIMITER, config.RedisPort)

	pool := redis.Pool{
		MaxIdle:     config.MaxIdle,
		IdleTimeout: config.IdleTimeout,
		MaxActive:   config.MaxActive,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(
				config.Protocol,
				redisAddress,
			)

			if err != nil {
				return nil, err
			}

			return conn, nil
		},
	}

	return &pool, nil
}
