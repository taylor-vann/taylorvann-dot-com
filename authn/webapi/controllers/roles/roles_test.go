package roles

import (
	"testing"
)

var testRole = CreateParams{
	UserID: 1,
	Role:   "internal",
}

var testRoleRead = ReadParams{
	UserID: 1,
	Role:   "internal",
}

func TestCreateTable(t *testing.T) {
	results, err := CreateTable()
	if err != nil {
		t.Error("error creating table.")
	}
	if results == nil {
		t.Error("no results were returned from CreateTable.")
	}
}

func TestCreateRow(t *testing.T) {
	result, err := Create(&testRole)
	if err != nil {
		t.Error("error creating Roles row.")
	}
	if result == nil {
		t.Error("no results were returned from Create.")
	}
	if result != nil && result.UserID != testRole.UserID {
		t.Error("failed to create Roles row.")
	}
}

func TestReadRow(t *testing.T) {
	result, err := Read(&testRoleRead)
	if err != nil {
		t.Error("error creating roles row.")
	}
	if result == nil {
		t.Error("no results were returned from Create.")
		return
	}
	if result.UserID != testRole.UserID {
		t.Error("mutated UserID.")
	}
	if result.Role != testRole.Role {
		t.Error("failed to read roles.")
	}
}

func TestRemoveRow(t *testing.T) {
	result, err := Remove(&testRoleRead)
	if err != nil {
		t.Error("error removing roles row.")
	}
	if result == nil {
		t.Error("No results were returned from Remove.")
		return
	}
	if result.UserID != testRole.UserID {
		t.Error("Mutated UserID.")
	}
	if result.Role != testRole.Role {
		t.Error("Deleted Role does not match correct Role.")
	}
}
