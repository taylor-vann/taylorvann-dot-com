package whitelistx

import (
	"encoding/json"
	"testing"
	"time"
)

type WhitelistxPayload struct {
	HelloWorld string `json:"hello_world"`
}

func TestPing(t *testing.T) {
	result, errPing := Ping()
	if errPing != nil {
		t.Error("Ping resulted in an error")
	}

	if *result != "PONG" {
		t.Error("Failed to Pong")
	}
}

func TestSetAndExpire(t *testing.T) {
	redisKey := "hello_world"
	payload := WhitelistxPayload{
		HelloWorld: "What's good, yo",
	}

	marshalledPayload, errMarshal := json.Marshal(payload)
	if errMarshal != nil {
		t.Error("error marshalling demo payload")
	}
	result, errResult := SetAndExpire(
		redisKey,
		&marshalledPayload,
		250,
	)
	if errResult != nil {
		t.Error("error setting and expiring")
	}

	if result == false {
		t.Error("could not set payload")
	}

	time.Sleep(251 * time.Millisecond)

	resultGet, errGet := Get(redisKey)
	if errGet != nil {
		t.Error("error getting expired key")
	}
	if resultGet != nil {
		t.Error("supposed to return nil")
	}
}

func TestGet(t *testing.T) {
	redisKey := "hello_world"
	payload := WhitelistxPayload{
		HelloWorld: "What's good, yo",
	}

	marshalledPayload, errMarshal := json.Marshal(payload)
	if errMarshal != nil {
		t.Error("error marshalling demo payload")
	}
	result, errResult := SetAndExpire(
		redisKey,
		&marshalledPayload,
		1000,
	)
	if errResult != nil {
		t.Error("error setting and expiring")
	}

	if result == false {
		t.Error("could not set payload")
	}

	resultGet, errResultGet := Get(redisKey)
	if errResultGet != nil {
		t.Error("error getting key value")
	}

	var whitelistPayload WhitelistxPayload
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
	payload := WhitelistxPayload{
		HelloWorld: "What's good, yo",
	}

	marshalledPayload, errMarshal := json.Marshal(payload)
	if errMarshal != nil {
		t.Error("error marshalling demo payload")
	}
	result, errResult := SetAndExpire(
		redisKey,
		&marshalledPayload,
		1000,
	)
	if errResult != nil {
		t.Error("error setting and expiring")
	}

	if result == false {
		t.Error("could not set payload")
	}

	resultGet, errResultGet := Get(redisKey)
	if errResultGet != nil {
		t.Error("error getting key value")
	}

	var whitelistPayload WhitelistxPayload
	errUnmarshal := json.Unmarshal(*resultGet, &whitelistPayload)
	if errUnmarshal != nil {
		t.Error("error unmarshalling payload")
	}

	if payload.HelloWorld != whitelistPayload.HelloWorld {
		t.Error("mismatching payloads")
	}

	resultRemove, errRemove := Remove(redisKey)
	if errRemove != nil {
		t.Error("error removing key")
	}
	if resultRemove == false {
		t.Error("could not remove key")
	}
}
