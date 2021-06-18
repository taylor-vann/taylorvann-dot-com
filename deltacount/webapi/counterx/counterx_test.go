package counterx

import (
	"testing"

	"webapi/details"
)

const (
	testUser = "TEST_USER"
)

func TestCreate(t *testing.T) {
	counter, errCounter := Create(&details.Details.Cache)
	if errCounter != nil {
		t.Fail()
		t.Logf(errCounter.Error())
	}

	if counter == nil {
		t.Fail()
		t.Logf("counter should not be nil")
	}
}

func TestIncrementAndGetCount(t *testing.T) {
	counter, errCounter := Create(&details.Details.Cache)
	if errCounter != nil {
		t.Fail()
		t.Logf(errCounter.Error())
	}

	count, errCount := counter.Increment(testUser)
	if count < 1 {
		t.Fail()
		t.Logf("counter.Increment should retrun a number larger than 0")
	}

	if errCount != nil {
		t.Fail()
		t.Logf(errCount.Error())
	}

	foundCount, errFoundCount := counter.Get(testUser)
	if errFoundCount != nil {
		t.Fail()
		t.Logf(errFoundCount.Error())
	}

	if count != foundCount {
		t.Fail()
		t.Logf("count should equal found count")
	}
}
