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
	Host        string				`json:"redis_host"`
	IdleTimeout time.Duration	`json:"idle_timeout"`
	MaxActive   int						`json:"max_active"`
	MaxIdle     int						`json:"max_idle"`
	MaxSizeInMB	string				`json:"max_size_in_mb"`
	Port        int						`json:"redis_port"`
	Protocol    string				`json:"protocol"`
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

	redisAddress := fmt.Sprint(config.Host, DELIMITER, config.Port)

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
