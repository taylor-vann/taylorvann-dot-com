package setterx

import (
	"testing"

	"webapi/details"

	"github.com/gomodule/redigo/redis"
)

const (
	testEntry  = "TEST_ENTRY"
	testResult = "TEST_RESULT"
)

func TestCreate(t *testing.T) {
	setter, errSetter := Create(&details.Details.Cache)
	if errSetter != nil {
		t.Fail()
		t.Logf(errSetter.Error())
	}

	if setter == nil {
		t.Fail()
		t.Logf("setter should not be nil")
	}
}

func TestGetAndSetCount(t *testing.T) {
	setter, errSetter := Create(&details.Details.Cache)
	if errSetter != nil {
		t.Fail()
		t.Logf(errSetter.Error())
		return
	}

	entry, errEntry := setter.Set(&SetBody{
		Address: testEntry,
		Entry:   testResult,
	})
	if errEntry != nil {
		t.Fail()
		t.Logf(errEntry.Error())
	}

	if entry == nil {
		t.Fail()
		t.Logf("setter.Set should retrun an entry")
	}

	getterEntry, errGetterEntry := redis.String(setter.Get(testEntry))
	if errGetterEntry != nil {
		t.Fail()
		t.Logf(errGetterEntry.Error())
	}

	if getterEntry != testResult {
		t.Fail()
		t.Logf("setter.Get should equal found count")
	}

}
