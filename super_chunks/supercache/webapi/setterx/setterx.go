package setterx

import (
	"webapi/redisx"

	"github.com/gomodule/redigo/redis"
)

type Setter struct {
	Store *redis.Pool `json:"store"`
}

type SetBody struct {
	Address string      `json:"address"`
	Entry   interface{} `json:"entry"`
}

func (s *Setter) Set(rBody *SetBody) (interface{}, error) {
	conn := s.Store.Get()
	defer conn.Close()

	return conn.Do("SET", rBody.Address, rBody.Entry)
}

func (s *Setter) Get(key string) (interface{}, error) {
	conn := s.Store.Get()
	defer conn.Close()

	return conn.Do("GET", key)
}

func Create(config *redisx.Config) (*Setter, error) {
	store, err := redisx.Create(config)
	return &Setter{store}, err
}
