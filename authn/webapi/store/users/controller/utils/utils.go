// brian taylor vann
// taylorvann dot com

package utils

import "time"

type MilliSeconds = int64

func GetNowAsMS() MilliSeconds {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
