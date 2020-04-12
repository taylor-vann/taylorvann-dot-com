package users

import (
	"testing"
)

var testUser = CreateParams{
	Email:    "test_user@test_email.test",
	Password: "pazzw0rd",
}

var testUserTwo = CreateParams{
	Email:    "test_user2@test_email.test",
	Password: "pazzw0rd",
}

var testUserRead = ReadParams{
	Email: "test_user@test_email.test",
}

var testUserSearch = SearchParams{
	EmailSubstring: "test_email",
}

var testUserUpdated = UpdateParams{
	CurrentEmail: "test_user@test_email.test",
	UpdatedEmail: "next_test_user@test_email.test",
	Password: "pazzword",
	IsDeleted: false,
}

var testUserUpdatedEmail = UpdateEmailParams {
	CurrentEmail: "next_test_user@test_email.test",
	UpdatedEmail: "another_test_user@test_email.test",
}

var testUserUpdatedPassword = UpdatePasswordParams {
	Email: "another_test_user@test_email.test",
	Password: "pazzw3rd",
}

var testUserRemoveUpdated = RemoveParams{
	Email: "another_test_user@test_email.test",
}

var testUserReviveUpdated = ReviveParams{
	Email: "another_test_user@test_email.test",
}

func TestCreateTestTable(t *testing.T) {
	results, err := CreateTable()
	if err != nil {
		t.Error("Error creating test table.")
	}
	if results == nil {
		t.Error("No results were returned from CreateTable.")
	}
}

func TestCreateRow(t *testing.T) {
	rows, err := Create(&testUser)
	if err != nil {
		t.Error("Error creating user row.")
		return
	}
	if len(rows) > 1 || len(rows) == 0 {
		t.Error("No results were returned from Update.")
		return
	}

	result := rows[0]

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
	rows, err := Read(&testUserRead)
	if err != nil {
		t.Error("Error reading user row.")
		return
	}
	if len(rows) > 1 || len(rows) == 0 {
		t.Error("No results were returned from Update.")
		return
	}

	result := rows[0]

	if result.Email != testUser.Email {
		t.Error("Failed to read user row.")
	}
	if result.CreatedAt != result.UpdatedAt {
		t.Error("CreatedAt does not equal UpdatedAt.")
	}
}

func TestSearchRows(t *testing.T) {
	newUserRow, err := Create(&testUserTwo)
	if err != nil {
		t.Error("Error creating user row.")
		return
	}
	if len(newUserRow) == 0 {
		t.Error("No results were returned from creating user.")
		return
	}

	rows, err := Search(&testUserSearch)
	if err != nil {
		t.Error("Error reading user row.")
		return
	}
	if len(rows) == 0 {
		t.Error("No results were returned from Update.")
		return
	}

	if len(rows) < 2 {
		t.Error("More than two results were supposed to be returned")
		return
	}
}

func TestUpdateRow(t *testing.T) {
	rows, err := Update(&testUserUpdated)
	if err != nil {
		t.Error("Error updating user row.")
		return
	}
	if len(rows) > 1 || len(rows) == 0 {
		t.Error("No results were returned from Update.")
		return
	}

	result := rows[0]

	if result.Email == testUser.Email {
		t.Error("Failed to updated email.")
	}
	if result.Email != testUserUpdated.UpdatedEmail {
		t.Error("Failed to correctly update email.")
	}
	if result.CreatedAt == result.UpdatedAt {
		t.Error("CreatedAt should not equal UpdatedAt.")
	}
}

func TestUpdateEmail(t *testing.T) {
	rows, err := UpdateEmail(&testUserUpdatedEmail)
	if err != nil {
		t.Error("Error updating user row.")
		return
	}
	if len(rows) > 1 || len(rows) == 0 {
		t.Error("No results were returned from Update.")
		return
	}

	result := rows[0]

	if result.Email == testUser.Email {
		t.Error("Failed to updated email.")
	}
	if result.Email != testUserUpdatedEmail.UpdatedEmail {
		t.Error("Failed to correctly update email.")
	}
	if result.CreatedAt == result.UpdatedAt {
		t.Error("CreatedAt should not equal UpdatedAt.")
	}
}

func TestUpdatePassword(t *testing.T) {
	rows, err := UpdatePassword(&testUserUpdatedPassword)
	if err != nil {
		t.Error("Error updating user row.")
		return
	}
	if len(rows) > 1 || len(rows) == 0 {
		t.Error("No results were returned from Update.")
		return
	}

	result := rows[0]

	if result.Email != testUserUpdatedPassword.Email {
		t.Error("Failed to update correct account password.")
	}
	if result.CreatedAt == result.UpdatedAt {
		t.Error("CreatedAt should not equal UpdatedAt.")
	}
}

func TestRemoveRow(t *testing.T) {
	rows, err := Remove(&testUserRemoveUpdated)
	if err != nil {
		t.Error("Error removing user row.")
		return
	}
	if len(rows) == 0 {
		t.Error("No results were returned from removal.")
		return
	}

	result := rows[0]

	if result.Email != testUserRemoveUpdated.Email {
		t.Error("Failed to remove correct user.")
	}
	if result.IsDeleted == false {
		t.Error("IsDeleted should be true")
	}
}

func TestReviveRow(t *testing.T) {
	rows, err := Revive(&testUserReviveUpdated)
	if err != nil {
		t.Error("Error reviving user row.")
		return
	}
	if len(rows) == 0 {
		t.Error("No results were returned from removal.")
		return
	}

	result := rows[0]

	if result.Email != testUserReviveUpdated.Email {
		t.Error("Failed to remove correct user.")
	}
	if result.IsDeleted == true {
		t.Error("IsDeleted should be false")
	}
}
