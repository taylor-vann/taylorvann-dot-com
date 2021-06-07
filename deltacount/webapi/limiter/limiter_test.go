package limiter

import (
	"fmt"
	"testing"
)

func TestGetIntervals(t *testing.T) {
	prevInterval, currInterval := getIntervals()

	adjPrevInterval := (prevInterval + 1) % adjDayAsSeconds
	if adjPrevInterval != currInterval {
		t.Fail()
		t.Logf(fmt.Sprint("intervals should be 1 int apart: ", prevInterval, currInterval))
	}
}
