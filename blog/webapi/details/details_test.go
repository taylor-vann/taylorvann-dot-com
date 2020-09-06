package details

import (
	"testing"
)

func TestReadDetailsFromFile(t *testing.T) {
	details := ReadDetailsFromFile()

	if details == nil {
		t.Error("server details are nil")
	}
}
