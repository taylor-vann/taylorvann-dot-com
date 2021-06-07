package counterx

import (
	"webapi/details"
	"webapi/redisx"

	"github.com/gomodule/redigo/redis"
)

type Config struct {
	Limiter   details.LimiterDetails `json:"limiter"`
	BlockList []string               `json:"block_list"`
	Cache     redisx.Config          `json:"cache"`
}

type Counter struct {
	Store *redis.Pool
}

func (c *Counter) Get(key string) (int, error) {
	conn := c.Store.Get()
	defer conn.Close()

	return redis.Int(conn.Do("GET", key))
}

func (c *Counter) Increment(key string) (int, error) {
	conn := c.Store.Get()
	defer conn.Close()

	return redis.Int(conn.Do("INCR", key))
}

func Create(config *redisx.Config) (*Counter, error) {
	store, errStore := redisx.Create(config)
	if errStore != nil {
		return nil, errStore
	}

	counter := Counter{
		Store: store,
	}

	return &counter, nil
}
