package haspassword

import (
	"fmt"
	"testing"
)

var testHasPassword = CreateParams{
	UserID: 1,
	PasswordID: 3,
}

var testHasPasswordRead = ReadParams{
	UserID: 1,
}

var testHasPasswordUpdate = CreateParams{
	UserID: 1,
	PasswordID: 7,
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
	result, err := Create(&testHasPassword)
	if err != nil {
		t.Error("Error creating haspassword row.")
		fmt.Println(err)
	}
	if result == nil {
		t.Error("No results were returned from Create.")
	}
	if result != nil && result.UserID != testHasPassword.UserID {
		t.Error("Failed to create haspassword row.")
	}
	if result != nil && result.CreatedAt != result.UpdatedAt {
		t.Error("CreatedAt does not equal UpdatedAt.")
	}
}

func TestReadRow(t *testing.T) {
	result, err := Read(&testHasPasswordRead)
	if err != nil {
		t.Error("Error creating haspassword row.")
		fmt.Println(err)
	}
	if result == nil {
		t.Error("No results were returned from Create.")
	}
	if result != nil && result.UserID != testHasPassword.UserID {
		t.Error("Mutated UserID.")
	}
	if result != nil && result.PasswordID != testHasPassword.PasswordID {
		t.Error("Failed to updated haspassword.")
	}
}

func TestUpdateRow(t *testing.T) {
	result, err := Update(&testHasPasswordUpdate)
	if err != nil {
		t.Error("Error updating haspassword row.")
		fmt.Println(err)
	}
	if result == nil {
		t.Error("No results were returned from Update.")
	}
	if result != nil && result.UserID != testHasPasswordUpdate.UserID {
		t.Error("Mutated UserID.")
	}
	if result != nil && result.PasswordID == testHasPassword.PasswordID {
		t.Error("Failed to updated haspassword.")
	}
	if result != nil && result.CreatedAt == result.UpdatedAt {
		t.Error("CreatedAt should not equal UpdatedAt.")
	}
}


func TestRemoveRow(t *testing.T) {
	result, err := Remove(&testHasPasswordRead)
	if err != nil {
		t.Error("Error removing haspassword row.")
		fmt.Println(err)
	}
	if result == nil {
		t.Error("No results were returned from Remove.")
	}
	if result != nil && result.UserID != testHasPasswordUpdate.UserID {
		t.Error("Mutated UserID.")
	}
	if result != nil && result.PasswordID != testHasPasswordUpdate.PasswordID {
		t.Error("Deleted PasswordID does not match correct PasswordID.")
	}
}
