// brian taylor vann
// taylorvann-dot-com

package controller

import (
	"testing"
)

var createTable = CreateTableParams{
	Environment: "LOCAL",
}

var testUser = CreateParams{
	Environment: "LOCAL",
	Email:    "test_user@test_email.test",
	Password: "pazzw0rd",
}

var testUserTwo = CreateParams{
	Environment: "LOCAL",
	Email:    "test_user2@test_email.test",
	Password: "pazzw0rd",
}

var testUserRead = ReadParams{
	Environment: "LOCAL",
	Email: "test_user@test_email.test",
}

var testUserSearch = SearchParams{
	Environment: "LOCAL",
	EmailSubstring: "test_email",
}

var testUserIndex = IndexParams{
	Environment: "LOCAL",
	StartIndex: 0,
	Length: 1024,
}

var testUserUpdated = UpdateParams{
	Environment: "LOCAL",
	CurrentEmail: "test_user@test_email.test",
	UpdatedEmail: "complete_update_test_user@test_email.test",
	Password: "pazzword",
	IsDeleted: false,
}

var testUserUpdatedEmail = UpdateEmailParams {
	Environment: "LOCAL",
	CurrentEmail: "complete_update_test_user@test_email.test",
	UpdatedEmail: "updated_email_test_user@test_email.test",
}

var testUserUpdatedPassword = UpdatePasswordParams {
	Environment: "LOCAL",
	Email: "updated_email_test_user@test_email.test",
	Password: "pazzw3rd",
}

var testUserRemoveUpdated = DeleteParams{
	Environment: "LOCAL",
	Email: "updated_email_test_user@test_email.test",
}

var testUserReviveUpdated = UndeleteParams{
	Environment: "LOCAL",
	Email: "updated_email_test_user@test_email.test",
}

func TestCreateTestTable(t *testing.T) {
	results, err := CreateTable(&createTable)
	if err != nil {
		t.Error(err.Error())
	}
	if results == nil {
		t.Error("No results were returned from CreateTable.")
	}
}

func TestCreate(t *testing.T) {
	rows, err := Create(&testUser)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if len(rows) == 0 {
		t.Error("No results were returned from Create.")
		return
	}
	if len(rows) != 1 {
		t.Error("Incorrect amount of results were returned from Create.")
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

func TestRead(t *testing.T) {
	rows, err := Read(&testUserRead)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if len(rows) == 0 {
		t.Error("No results were returned from Read.")
		return
	}
	if  len(rows) != 1 {
		t.Error("Incorrect amount of rows were returned from Read.")
		return
	}

	result := rows[0]

	if result.Email != testUser.Email {
		t.Error("Failed to read user row.")
	}
}

func TestValidate(t *testing.T) {
	rows, err := Validate(&testUser)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if len(rows) == 0 {
		t.Error("No results were returned from Validate.")
		return
	}
	if  len(rows) != 1 {
		t.Error("Incorrect amount of rows were returned from Validate.")
		return
	}

	result := rows[0]

	if result.Email != testUser.Email {
		t.Error("Failed to read user row.")
	}
}

func TestValidateBadPassword(t *testing.T) {
	rows, err := Validate(&ValidateParams{
		Email: testUser.Email,
		Password: "kawabunga!",
	})
	if err != nil {
		t.Error(err.Error())
		return
	}
	if len(rows) != 0 {
		t.Error("results were returned from bad password.")
		return
	}
	if len(rows) == 1 {
		t.Error("incorrect amount of rows were returned from bad password.")
		return
	}
}

func TestSearch(t *testing.T) {
	newUserRow, err := Create(&testUserTwo)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if len(newUserRow) == 0 {
		t.Error("No results were returned from creating user.")
		return
	}

	rows, err := Search(&testUserSearch)
	if err != nil {
		t.Error(err.Error())
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

func TestIndex(t *testing.T) {
	rows, err := Index(&testUserIndex)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if len(rows) == 0 {
		t.Error("No results were returned from Index.")
		return
	}

	if len(rows) < 2 {
		t.Error("More than two results were supposed to be returned")
		return
	}
}

func TestUpdate(t *testing.T) {
	rows, err := Update(&testUserUpdated)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if len(rows) == 0 {
		t.Error("No results were returned from Update.")
		return
	}
	if  len(rows) != 1 {
		t.Error("Incorrect amount of rows were returned from Read.")
		return
	}

	result := rows[0]

	if result.Email == testUser.Email {
		t.Error("Failed to updated email.")
	}
	if result.Email != testUserUpdated.UpdatedEmail {
		t.Error("Failed to correctly update user.")
	}
	if result.CreatedAt == result.UpdatedAt {
		t.Error("CreatedAt should not equal UpdatedAt.")
	}
}

func TestUpdateEmail(t *testing.T) {
	rows, err := UpdateEmail(&testUserUpdatedEmail)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if len(rows) == 0 {
		t.Error("No results were returned from UpdateEmail.")
		return
	}
	if  len(rows) != 1 {
		t.Error("Incorrect amount of rows were returned from Read.")
		return
	}

	result := rows[0]

	if result.Email == testUser.Email {
		t.Error("Failed to update email.")
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
		t.Error(err.Error())
		return
	}
	if len(rows) == 0 {
		t.Error("No results were returned from UpdatePassword.")
		return
	}
	if  len(rows) != 1 {
		t.Error("Incorrect amount of rows were returned from UpdatePassword.")
		return
	}

	result := rows[0]

	if result.Email != testUserUpdatedPassword.Email {
		t.Error("Failed to update password.")
	}
	if result.CreatedAt == result.UpdatedAt {
		t.Error("CreatedAt should not equal UpdatedAt.")
	}
}

func TestDelete(t *testing.T) {
	rows, err := Delete(&testUserRemoveUpdated)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if len(rows) == 0 {
		t.Error("No results were returned from Remove.")
		return
	}
	if len(rows) != 1 {
		t.Error("Incorrect amount of rows were returned from Remove.")
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

func TestUndelete(t *testing.T) {
	rows, err := Undelete(&testUserReviveUpdated)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if len(rows) == 0 {
		t.Error("No results were returned from Revive.")
		return
	}
	if len(rows) != 1 {
		t.Error("Incorrect amount of rows were returned from Revive.")
		return
	}

	result := rows[0]

	if result.Email != testUserReviveUpdated.Email {
		t.Error("Failed to revive correct user.")
	}
	if result.IsDeleted == true {
		t.Error("IsDeleted should be false")
	}
}

func TestDangerouslyDropUnitTestsTable(t *testing.T) {
	result, err := DangerouslyDropUnitTestsTable()

	if result == nil {
		t.Error("Failed to drop table")
	}
	if err != nil {
		t.Error(err.Error())
	}
}