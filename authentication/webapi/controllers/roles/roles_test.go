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
}

var testRoleUpdate = CreateParams{
	UserID: 1,
	Role:   "admin",
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
	if result != nil && result.CreatedAt != result.UpdatedAt {
		t.Error("CreatedAt does not equal UpdatedAt.")
	}
}

func TestReadRow(t *testing.T) {
	result, err := Read(&testRoleRead)
	if err != nil {
		t.Error("error creating roles row.")
	}
	if result == nil {
		t.Error("no results were returned from Create.")
	}
	if result != nil && result.UserID != testRole.UserID {
		t.Error("mutated UserID.")
	}
	if result != nil && result.Role != testRole.Role {
		t.Error("failed to updated roles.")
	}
}

func TestUpdateRow(t *testing.T) {
	result, err := Update(&testRoleUpdate)
	if err != nil {
		t.Error("error updating roles row.")
	}
	if result == nil {
		t.Error("no results were returned from Update.")
	}
	if result != nil && result.UserID != testRoleUpdate.UserID {
		t.Error("mutated UserID.")
	}
	if result != nil && result.Role == testRole.Role {
		t.Error("failed to updated roles.")
	}
	if result != nil && result.CreatedAt == result.UpdatedAt {
		t.Error("CreatedAt should not equal UpdatedAt.")
	}
}

func TestRemoveRow(t *testing.T) {
	result, err := Remove(&testRoleRead)
	if err != nil {
		t.Error("error removing roles row.")
	}
	if result == nil {
		t.Error("No results were returned from Remove.")
	}
	if result != nil && result.UserID != testRoleUpdate.UserID {
		t.Error("Mutated UserID.")
	}
	if result != nil && result.Role != testRoleUpdate.Role {
		t.Error("Deleted Role does not match correct Role.")
	}
}
