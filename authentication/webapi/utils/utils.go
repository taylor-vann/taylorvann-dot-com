package utils

import "time"

// MilliSecond -
type MilliSecond = int64

// GetNowAsMS -
func GetNowAsMS() MilliSecond {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
