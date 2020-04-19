// brian taylor vann
// taylorvann dot com

package graylistx

import (
	"errors"
	"github.com/gomodule/redigo/redis"

	"webapi/interfaces/redisx"
	"webapi/interfaces/graylistx/constants"
)

type MilliSeconds = int64

type Config = redisx.RedisConfig

type SetAndExpireParams struct {
	Key 				string
	Value 			[]byte
	ExpiryInMS 	MilliSeconds
}

type GetParams struct {
	Key 				string
}

type RemoveParams = GetParams

type Instance struct {
	Connection *redisx.Connection
}

func (instance *Instance) Ping() (*string, error) {
	if instance.Connection.Store == nil {
		return nil, errors.New("instance store is nil")
	}

	conn, errConn := instance.Connection.Store.Dial()
	if errConn != nil {
		return nil, errConn
	}

	defer conn.Close()
	result, errResult := redis.String(conn.Do("PING"))
	if errResult != nil {
		return nil, errResult
	}

	return &result, nil
}

func (instance *Instance) SetAndExpire(p *SetAndExpireParams) (bool, error) {
	if instance.Connection.Store == nil {
		return false, errors.New("instance store is nil")
	}

	conn, errConn := instance.Connection.Store.Dial()
	if errConn != nil {
		return false, errConn
	}

	defer conn.Close()
	result, errResult := redis.String(
		conn.Do(
			constants.Set,
			p.Key,
			p.Value,
			constants.Px,
			p.ExpiryInMS,
		),
	)
	if errResult != nil {
		return false, errResult
	}

	return result == constants.Ok, errResult
}

func (instance *Instance) Get(p *GetParams) (*[]byte, error) {
	if instance.Connection.Store == nil {
		return nil, errors.New("instance store is nil")
	}

	conn, errConn := instance.Connection.Store.Dial()
	if errConn != nil {
		return nil, errConn
	}

	defer conn.Close()
	result, errResult := conn.Do(constants.Get, p.Key)
	if errResult != nil {
		return nil, errResult
	}

	if result == nil {
		return nil, errResult
	}

	resultAsBytes, errResultAsBytes := redis.Bytes(result, errResult)
	if errResultAsBytes != nil {
		return nil, errResultAsBytes
	}

	return &resultAsBytes, nil
}

func (instance *Instance) Remove(p *RemoveParams) (bool, error) {
	if instance.Connection.Store == nil {
		return false, errors.New("redisInst is nil")
	}

	conn, errConn := instance.Connection.Store.Dial()
	if errConn != nil {
		return false, errConn
	}
	
	defer conn.Close()
	result, errResult := redis.Int(
		conn.Do(constants.Del, p.Key),
	)

	return result == 1, errResult
}

func Create(config Config) (*Instance, error) {
	connection, errConnection := redisx.Create(&config)
	if errConnection != nil {
		return nil, errConnection
	}
	instance := Instance{
		Connection: connection,
	}
	return &instance, nil
}