package utils

import (
	"math/rand"
	"time"
)

// MilliSeconds -
type MilliSeconds = int64

// GetNowAsMS -
func GetNowAsMS() MilliSeconds {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// GenerateRandomByteArray -
func GenerateRandomByteArray(n uint32) (*[]byte, error) {
	token := make([]byte, n)
	_, err := rand.Read(token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}
