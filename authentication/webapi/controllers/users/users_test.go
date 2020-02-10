package users

import (
	"fmt"
	"testing"
)

var testUser = CreateParams{
	Email: "test_user@test_email.test",
}

var testUserUpdated = UpdateParams{
	CurrentEmail: "test_user@test_email.test",
	UpdatedEmail: "next_test_user@test_email.test",
}

var testUserRemoveUpdated = CreateParams{
	Email: "next_test_user@test_email.test",
}

func TestCreateTable(t *testing.T) {
	results, err := CreateTable()
	if err != nil {
		t.Error("Error creating table.")
	}
	if results == nil {
		t.Error("No results were returned from CreateTable.")
	}
}

func TestCreateRow(t *testing.T) {
	result, err := Create(&testUser)
	if err != nil {
		t.Error("Error creating user row.")
		fmt.Println(err)
	}

	if result == nil {
		t.Error("No results were returned from Create.")
	}
	if result != nil && result.Email != testUser.Email {
		t.Error("Failed to create user row.")
	}
	if result != nil && result.IsDeleted == true {
		t.Error("Row should be deleted")
	}
	if result != nil && result.CreatedAt != result.UpdatedAt {
		t.Error("CreatedAt does not equal UpdatedAt.")
	}
}

func TestReadRow(t *testing.T) {
	result, err := Read(&testUser)
	if err != nil {
		t.Error("Error reading user row.")
		fmt.Println(err)
	}
	if result == nil {
		t.Error("No results were returned from Read.")
	}
	if result != nil && result.Email != testUser.Email {
		t.Error("Failed to read user row.")
	}
	if result != nil && result.CreatedAt != result.UpdatedAt {
		t.Error("CreatedAt does not equal UpdatedAt.")
	}
}

func TestUpdateRow(t *testing.T) {
	result, err := Update(&testUserUpdated)
	if err != nil {
		t.Error("Error updating user row.")
		fmt.Println(err)
	}
	if result == nil {
		t.Error("No results were returned from Update.")
	}
	if result != nil && result.Email == testUser.Email {
		t.Error("Failed to updated email.")
	}
	if result != nil && result.Email != testUserUpdated.UpdatedEmail {
		t.Error("Failed to correctly update email.")
	}
	if result != nil && result.CreatedAt == result.UpdatedAt {
		t.Error("CreatedAt should not equal UpdatedAt.")
	}
}

func TestRemoveRow(t *testing.T) {
	result, err := Remove(&testUserRemoveUpdated)
	if err != nil {
		t.Error("Error removing user row.")
		fmt.Println(err)
	}
	if result == nil {
		t.Error("No results were returned from Remove.")
	}
	if result != nil && result.Email != testUserUpdated.UpdatedEmail {
		t.Error("Failed to remove correct user.")
	}
	if result != nil && result.IsDeleted == false {
		t.Error("IsDeleted should be true")
	}
}
