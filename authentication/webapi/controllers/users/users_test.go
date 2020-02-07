package users

import (
	"fmt"
	"testing"
)

var testUser = CreateParams{
	"test_user@test_email.test",
}

var testUserUpdated = CreateParams{
	"next_test_user@test_email.test",
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
	if result.Email != testUser.Email {
		t.Error("Failed to create user row.")
	}
	if result.IsDeleted == true {
		t.Error("Row should be deleted")
	}
	if result.CreatedAt != result.UpdatedAt {
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
	if result.Email != testUser.Email {
		t.Error("Failed to read user row.")
	}
	if result.CreatedAt != result.UpdatedAt {
		t.Error("CreatedAt does not equal UpdatedAt.")
	}
	fmt.Println(result)
}

func TestUpdateRow(t *testing.T) {
	result, err := Update(&testUser, &testUserUpdated)
	if err != nil {
		t.Error("Error updating user row.")
		fmt.Println(err)
	}
	if result == nil {
		t.Error("No results were returned from Update.")
	}
	if result.Email == testUser.Email {
		t.Error("Failed to updated user.")
	}
	if result.Email != testUserUpdated.Email {
		t.Error("Failed to correctly update user.")
	}
	if result.CreatedAt == result.UpdatedAt {
		t.Error("CreatedAt should not equal UpdatedAt.")
	}
}

func TestRemoveRow(t *testing.T) {
	result, err := Remove(&testUserUpdated)
	if err != nil {
		t.Error("Error removing user row.")
		fmt.Println(err)
	}
	if result == nil {
		t.Error("No results were returned from Remove.")
	}
	if result.Email == testUser.Email {
		t.Error("Failed to remove correct user.")
	}
	if result.CreatedAt == result.UpdatedAt {
		t.Error("CreatedAt should not equal UpdatedAt.")
	}
	if result.IsDeleted == false {
		t.Error("IsDeleted should be true")
	}
}
