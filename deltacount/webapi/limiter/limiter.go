package limiter

// get  interval over day from details
// ie what chunk of 15 sec or whatever is it?

import (
	"time"
	"webapi/details"
	"webapi/counterx"
)

type BucketRequest struct {
	GlobalKey			string	`json:"global_key"`
	LocalKey			string	`json:"local_key"`
	Interval			int			`json:"interval"`
	PrevInterval	int			`json:"prev_interval"`
}

type BucketResposne struct {
	GlobalCount			int64	`json:"global_count"`
	PreviousGlobal	int64	`json:"previous_global_count"`
	LocalCount			int64	`json:"local_count"`
	PrevLocalCount	int64	`json:"prev_local_count"`
}

const (
	DROPPED_REQUESTS = "DROPPED_REQUESTS"
)

var (
	interval        = 1 / details.Details.Limiter.Interval
	adjDayAsSeconds = 60 * 60 * 24 * interval
	counter, errCounter = counterx.Create(&details.Details.Cache)
)

// in seconds
func getIntervals() (int, int) {
	now := time.Now()

	adjustedInterval := now.Hour()*60*60 + now.Minute()*60 + now.Second()*interval
	prevInterval := (adjustedInterval - 1) % adjDayAsSeconds

	return prevInterval, adjustedInterval
}

func GetDroppedRequests() (int, error) {
	if (errCounter) {
		return -1, errCounter
	}

	return counter.Increment(DROPPED_REQUESTS)
}

func incrementRequestCount(address string, err error) (int, error) {
	if (errCounter) {
		return -1, errCounter
	}

	if (err) {
		return -1, err
	}

	return counter.Increment(address)
}

func getRequestCount(address string, err error) (int, error) {
	if (errCounter) {
		return -1, errCounter
	}

	if (err) {
		return -1, err
	}

	return counter.Get(address)
}

func incrementDroppedRequests() (int, error) {
	return incrementRequestCount(DROPPED_REQUESTS)
}

func getDroppedRequests() (int, error) {
	return getRequestCount(DROPPED_REQUESTS)
}

// increment denied requests and return error
//

// global bucket check
// use previous interval
// measure ratio, if under 1 return or error
//

// user bucket check
// use previous interval
// measure ratio, if under 1 return or error

// increment global count
// get global count
// if over total limit per interval, (increment denied requests and return error)
// otherwise
// get prev minute and do bucket algorithm
// if does not pass bucket check return errors

// increment user count
// if over limit for user per interval, increment denied requests and return error
// otherwise get prev minute and do bucket
// if does not pass bucket check, return error
//

// hasBandwidth -> bool, error
// get interval
// global bucket check
// ip bucket check
