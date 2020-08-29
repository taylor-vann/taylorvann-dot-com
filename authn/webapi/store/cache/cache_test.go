package cache

import (
	"strconv"
	"testing"
	"time"
)

type TestEntry struct {
	CreatedAt int64
	Lifetime  int64
}

type TestPlan struct {
	Key   string
	Entry TestEntry
}

const TestEnvironment = "UNIT_TESTS"

var testPlans = generateRandomTestPlans(5)

func getNowAsMS() MilliSeconds {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func getLaterAsMS() MilliSeconds {
	return (time.Now().UnixNano() + DayAsMS) / int64(time.Millisecond)
}

func generateRandomTestPlans(num int) *[]TestPlan {
	testPlans := make([]TestPlan, num)

	for index := range testPlans {
		nowAsMS := getNowAsMS()

		indexAsStr := strconv.Itoa(index)
		nowAsMSAsStr := strconv.FormatInt(nowAsMS, 10)

		key := "UNIT_TEST_CACHE_ENTRY" + "_" + indexAsStr + "_" + nowAsMSAsStr
		testPlans[index] = TestPlan{
			Key: key,
			Entry: TestEntry{
				CreatedAt: nowAsMS,
				Lifetime:  DayAsMS,
			},
		}
	}

	return &testPlans
}

func TestCreateEntry(t *testing.T) {
	for _, entry := range *testPlans {
		entry, errEntry := CreateEntry(&CreateEntryParams{
			Environment: TestEnvironment,
			CreatedAt:   entry.Entry.CreatedAt,
			Lifetime:    entry.Entry.Lifetime,
			Payload:     entry,
			Key:         entry.Key,
		})

		if errEntry != nil {
			t.Error(errEntry.Error())
		}

		if entry == nil {
			t.Error("nil entry returned")
		}
	}
}

func TestReadEntry(t *testing.T) {
	var expected = make([]TestPlan, len(*testPlans))

	for index, test := range *testPlans {
		expected[index] = test
		entry, errEntry := CreateEntry(&CreateEntryParams{
			Environment: TestEnvironment,
			CreatedAt:   test.Entry.CreatedAt,
			Lifetime:    test.Entry.Lifetime,
			Payload:     test,
			Key:         test.Key,
		})

		if errEntry != nil {
			t.Error(errEntry.Error())
		}

		if entry == nil {
			t.Error("nil entry returned")
		}
	}

	// check entries for accuracy
	for _, test := range expected {
		readEntry, errReadEntry := ReadEntry(&ReadEntryParams{
			Environment: TestEnvironment,
			Key:         test.Key,
		})
		if errReadEntry != nil {
			t.Error(errReadEntry.Error())
		}
		if readEntry == nil {
			t.Error("entry should not be nil")
		}
		// check if entry matches
	}
}

func TestRemoveEntry(t *testing.T) {
	var results = make([]TestPlan, len(*testPlans))

	for index, test := range *testPlans {
		results[index] = test
		entry, errEntry := CreateEntry(&CreateEntryParams{
			Environment: TestEnvironment,
			CreatedAt:   test.Entry.CreatedAt,
			Lifetime:    test.Entry.Lifetime,
			Key:         test.Key,
			Payload:     test.Entry,
		})

		if errEntry != nil {
			t.Error(errEntry.Error())
		}

		if entry == nil {
			t.Error("nil entry returned")
		}
	}

	// check entries for accuracy
	for _, test := range results {
		removeEntry, errRemoveEntry := RemoveEntry(&RemoveEntryParams{
			Environment: TestEnvironment,
			Key:         test.Key,
		})
		if errRemoveEntry != nil {
			t.Error(errRemoveEntry.Error())
		}
		if removeEntry == false {
			t.Error("couldn't remove entry")
		}
	}
}
