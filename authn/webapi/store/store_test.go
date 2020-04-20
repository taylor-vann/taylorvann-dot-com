package store

import (
	"testing"
)

func TestCreateTables(t *testing.T) {
	_, errUsers := CreateRequiredTables()
	if errUsers != nil {
		t.Error("error creating tables from store")
	}
}