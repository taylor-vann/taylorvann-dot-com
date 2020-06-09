package graylistx

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"testing"
	"time"
)

type GraylistxPayload struct {
	HelloWorld string `json:"hello_world"`
}

const (
	whitelistHost               = "WHITELIST_HOST"
	whitelistPort               = "WHITELIST_PORT"
	whitelistProtocol           = "WHITELIST_PROTOCOL"
	whitelistMaxActive          = "WHITELIST_MAX_ACTIVE"
	whitelistMaxIdle            = "WHITELIST_MAX_IDLE"
	whitelistIdleTimeoutSeconds = "WHITELIST_IDLE_TIMEOUT_SECONDS"
)

var graylistConfig, errGraylistConfig = getWhitelistxEnvConstants()
var graylist, errGraylist = Create(graylistConfig)

func getWhitelistxEnvConstants() (*Config, error) {
	host := os.Getenv(whitelistHost)
	port := os.Getenv(whitelistPort)
	protocol := os.Getenv(whitelistProtocol)
	maxIdle := os.Getenv(whitelistMaxIdle)
	maxActive := os.Getenv(whitelistMaxActive)
	idleTimoutSeconds := os.Getenv(whitelistIdleTimeoutSeconds)

	if host == "" || port == "" || protocol == "" || maxIdle == "" || maxActive == "" || idleTimoutSeconds == "" {
		return nil, errors.New(
			"whitelistx - getConfigFromEnv - unable to import required evnironment variables",
		)
	}

	portAsInt, errPort := strconv.Atoi(port)
	if errPort != nil {
		return nil, errPort
	}

	maxIdleAsInt, errIdle := strconv.Atoi(maxIdle)
	if errIdle != nil {
		return nil, errIdle
	}

	maxActiveAsInt, errMaxActive := strconv.Atoi(maxActive)
	if errMaxActive != nil {
		return nil, errMaxActive
	}

	idleTimoutSecondsAsInt, errIdle := strconv.Atoi(idleTimoutSeconds)
	if errIdle != nil {
		return nil, errIdle
	}

	idleTimeoutAsTime := time.Duration(idleTimoutSecondsAsInt) * time.Second

	config := Config{
		Host:        host,
		Port:        portAsInt,
		Protocol:    protocol,
		MaxIdle:     maxIdleAsInt,
		IdleTimeout: idleTimeoutAsTime,
		MaxActive:   maxActiveAsInt,
	}

	return &config, nil
}

func TestPing(t *testing.T) {
	result, errPing := graylist.Ping()
	if errPing != nil {
		t.Error("Ping resulted in an error")
	}

	if *result != "PONG" {
		t.Error("Failed to Pong")
	}
}

func TestSetAndExpire(t *testing.T) {
	redisKey := "hello_world"
	payload := GraylistxPayload{
		HelloWorld: "What's good, yo",
	}

	marshalledPayload, errMarshal := json.Marshal(payload)
	if errMarshal != nil {
		t.Error("error marshalling demo payload")
	}
	result, errResult := graylist.SetAndExpire(&SetAndExpireParams{
		Key: redisKey,
		Value: marshalledPayload,
		ExpiryInMS: 250,
	})
	if errResult != nil {
		t.Error("error setting and expiring")
	}

	if result == false {
		t.Error("could not set payload")
	}

	time.Sleep(251 * time.Millisecond)

	resultGet, errGet := graylist.Get(&GetParams{
		Key: redisKey,
	})
	if errGet != nil {
		t.Error("error getting expired key")
	}
	if resultGet != nil {
		t.Error("supposed to return nil")
	}
}

func TestGet(t *testing.T) {
	redisKey := "hello_world"
	payload := GraylistxPayload{
		HelloWorld: "What's good, yo",
	}

	marshalledPayload, errMarshal := json.Marshal(payload)
	if errMarshal != nil {
		t.Error("error marshalling demo payload")
	}
	result, errResult := graylist.SetAndExpire(&SetAndExpireParams{
		Key: redisKey,
		Value: marshalledPayload,
		ExpiryInMS: 1000,
	})
	if errResult != nil {
		t.Error("error setting and expiring")
	}

	if result == false {
		t.Error("could not set payload")
	}

	resultGet, errResultGet := graylist.Get(&GetParams{
		Key: redisKey,
	})
	if errResultGet != nil {
		t.Error("error getting key value")
	}

	var whitelistPayload GraylistxPayload
	errUnmarshal := json.Unmarshal(*resultGet, &whitelistPayload)
	if errUnmarshal != nil {
		t.Error("error unmarshalling payload")
	}

	if payload.HelloWorld != whitelistPayload.HelloWorld {
		t.Error("mismatching payloads")
	}
}

func TestRemove(t *testing.T) {
	redisKey := "hello_world"
	payload := GraylistxPayload{
		HelloWorld: "What's good, yo",
	}

	marshalledPayload, errMarshal := json.Marshal(payload)
	if errMarshal != nil {
		t.Error("error marshalling demo payload")
	}
	result, errResult := graylist.SetAndExpire(&SetAndExpireParams{
		Key: redisKey,
		Value: marshalledPayload,
		ExpiryInMS: 1000,
	})
	if errResult != nil {
		t.Error("error setting and expiring")
	}

	if result == false {
		t.Error("could not set payload")
	}

	resultGet, errResultGet := graylist.Get(&GetParams{
		Key: redisKey,
	})
	if errResultGet != nil {
		t.Error("error getting key value")
	}

	var whitelistPayload GraylistxPayload
	errUnmarshal := json.Unmarshal(*resultGet, &whitelistPayload)
	if errUnmarshal != nil {
		t.Error("error unmarshalling payload")
	}

	if payload.HelloWorld != whitelistPayload.HelloWorld {
		t.Error("mismatching payloads")
	}

	resultRemove, errRemove := graylist.Remove(&RemoveParams{
		Key: redisKey,
	})
	if errRemove != nil {
		t.Error("error removing key")
	}
	if resultRemove == false {
		t.Error("could not remove key")
	}
}
