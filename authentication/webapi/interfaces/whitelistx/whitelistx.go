// brian taylor vann
// taylorvann dot com

package whitelistx

import (
	"errors"
	"github.com/gomodule/redigo/redis"

	"webapi/interfaces/redisx"
	"webapi/interfaces/whitelistx/constants"
	"webapi/utils"
)

// Get config from environemnt variables
func getConfigFromEnv() (*redisx.RedisConfig, error) {
	config := redisx.RedisConfig{
		Host:        constants.Env.Host,
		Port:        constants.Env.Port,
		Protocol:    constants.Env.Protocol,
		MaxIdle:     constants.Env.MaxIdle,
		IdleTimeout: constants.Env.IdleTimeout,
		MaxActive:   constants.Env.MaxActive,
	}

	return &config, nil
}

// get our config
var redisConf, errConfig = getConfigFromEnv()

// create instance of redis pool
var redisInst, errRedisx = redisx.Create(redisConf)

// Ping -
func Ping() (*string, error) {
	if redisInst == nil {
		return nil, errors.New("whitelistx - Ping - redisInst is nil")
	}

	conn, errConn := redisInst.Store.Dial()
	if errConn != nil {
		return nil, errConn
	}

	result, errResult := redis.String(conn.Do("PING"))
	conn.Close()
	if errResult != nil {
		return nil, errResult
	}

	return &result, nil
}

// SetAndExpire - controlled set function
func SetAndExpire(key string, value *[]byte, expireMS utils.MilliSeconds) (bool, error) {
	if redisInst == nil {
		return false, errors.New("whitelistx - Set - redisInst is nil")
	}

	conn, errConn := redisInst.Store.Dial()
	if errConn != nil {
		return false, errConn
	}

	result, errResult := redis.String(
		conn.Do(
			constants.Set,
			key,
			*value,
			constants.Px,
			expireMS,
		),
	)
	conn.Close()
	if errResult != nil {
		return false, errResult
	}

	return result == constants.Ok, errResult
}

// Get - controlled get function
func Get(key string) (*[]byte, error) {
	if redisInst == nil {
		return nil, errors.New("whitelistx - Set - redisInst is nil")
	}

	conn, errConn := redisInst.Store.Dial()
	if errConn != nil {
		return nil, errConn
	}

	result, errResult := conn.Do(constants.Get, key)
	if errResult != nil {
		return nil, errResult
	}

	if result == nil {
		return nil, errResult
	}

	resultAsBytes, errResultAsBytes := redis.Bytes(result, errResult)
	conn.Close()
	if errResultAsBytes != nil {
		return nil, errResultAsBytes
	}

	return &resultAsBytes, nil
}

// Remove - controlled set function
func Remove(key string) (bool, error) {
	if redisInst == nil {
		return false, errors.New("whitelistx - Set - redisInst is nil")
	}

	conn, errConn := redisInst.Store.Dial()
	if errConn != nil {
		return false, errConn
	}

	result, errResult := redis.Int(
		conn.Do(constants.Del, key),
	)
	conn.Close()
	if errResult != nil {
		return false, errResult
	}

	return result == 1, errResult
}
