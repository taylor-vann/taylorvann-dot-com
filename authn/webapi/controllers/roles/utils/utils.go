package utils

import "time"

// MilliSeconds -
type MilliSeconds = int64

// GetNowAsMS -
func GetNowAsMS() MilliSeconds {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
