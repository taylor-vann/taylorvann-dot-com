package constants

import "time"

type WhitelistxConstants struct {
	Host        string
	Port        int
	Protocol    string
	Expire      string
	MaxIdle     int
	IdleTimeout time.Duration
	MaxActive   int
}

const (
	Ok  = "OK"
	Set = "SET"
	Get = "GET"
	Del = "DEL"
	Px  = "PX"
)
