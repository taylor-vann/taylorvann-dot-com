package passwords

import (
	"fmt"
	"testing"

	"webapi/interfaces/passwordx"
)

var testPasswordID int64
var testPassword = CreateParams{
	Password: "pazzw0rd",
}
var testPasswordUpdated = CreateParams{
	Password: "razz13dazz13",
}

func verifyPassword(password string, salt string, hash string, p *HashParams) (bool, error) {
	hashResults := passwordx.HashResults{
		Salt: salt,
		Hash: hash,
		Params: p,
	}
	return passwordx.PasswordIsValid(password, &hashResults)
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
	result, err := Create(&testPassword)
	if err != nil {
		t.Error("Error creating password row.")
		fmt.Println(err)
	}
	if result == nil {
		t.Error("No results were returned from Create.")
	}

	passwordIsValid, errPassword := verifyPassword(
		testPassword.Password,
		result.Salt,
		result.Hash,
		result.Params,
	)
	if errPassword != nil {
		t.Error("Error verifying password")
	}
	if !passwordIsValid {
		t.Error("Password could not be verified")
	}

	if result.Hash == testPassword.Password {
		t.Error("Failed to create password row.")
	}
	if result.Params == nil {
		t.Error("Failed to unmarshal Params")
	}
	if result.CreatedAt != result.UpdatedAt {
		t.Error("CreatedAt does not equal UpdatedAt.")
	}

	testPasswordID = result.ID
}

func TestReadRow(t *testing.T) {
	result, err := Read(&ReadParams{ID: testPasswordID})
	if err != nil {
		t.Error("Error reading password row.")
	}
	if result == nil {
		t.Error("No results were returned from Create.")
	}

	passwordIsValid, errPassword := verifyPassword(
		testPassword.Password,
		result.Salt,
		result.Hash,
		result.Params,
	)
	if errPassword != nil {
		t.Error("Error verifying password")
	}
	if !passwordIsValid {
		t.Error("Password could not be verified")
	}

	if result.Hash == testPassword.Password {
		t.Error("Failed to create password row.")
	}
	if result.Params == nil {
		t.Error("Failed to unmarshal Params")
	}
}

func TestUpdateRow(t *testing.T) {
	readParams := ReadParams{
		ID: testPasswordID,
	}
	passwordResult, errReadPassword := Read(&readParams)
	if errReadPassword != nil {
		t.Error("error retrieving password to update")
	}

	passwordIsValid, errPassword := verifyPassword(
		testPassword.Password,
		passwordResult.Salt,
		passwordResult.Hash,
		passwordResult.Params,
	)
	if errPassword != nil {
		t.Error("Error verifying password")
	}
	if !passwordIsValid {
		t.Error("Password could not be verified")
	}

	newPasswordResult, errUpdate := Update(&readParams, &testPasswordUpdated)
	if errUpdate != nil {
		t.Error("Error updating password row.")
		fmt.Println(errUpdate)
	}

	newPasswordIsValid, errNewPassword := verifyPassword(
		testPasswordUpdated.Password,
		newPasswordResult.Salt,
		newPasswordResult.Hash,
		newPasswordResult.Params,
	)

	if errNewPassword != nil {
		t.Error("Error verifying new password")
	}
	if !newPasswordIsValid {
		t.Error("NewPassword could not be verified")
	}
	if passwordResult.ID != newPasswordResult.ID {
		t.Error("Old password and new password are not the same password row.")
	}
	if newPasswordResult == nil {
		t.Error("No results were returned from Create.")
	}
	if newPasswordResult.Hash == testPassword.Password {
		t.Error("Failed to create password row.")
	}
	if newPasswordResult.Params == nil {
		t.Error("Failed to unmarshal Params")
	}
}

func TestRemoveRow(t *testing.T) {
	readParams := ReadParams{
		ID: testPasswordID,
	}
	result, errRemove := Remove(&readParams)
	if errRemove != nil {
		t.Error("Error removing password row.")
		fmt.Println(errRemove)
	}

	passwordIsValid, errPassword := verifyPassword(
		testPasswordUpdated.Password,
		result.Salt,
		result.Hash,
		result.Params,
	)
	if errPassword != nil {
		t.Error("Error verifying password")
	}
	if !passwordIsValid {
		t.Error("Password could not be verified")
	}

	if result == nil {
		t.Error("No results were returned from Remove.")
	}
	if result.Hash == testPassword.Password {
		t.Error("Failed to hash password row.")
	}
	if result.Params == nil {
		t.Error("Failed to unmarshal Params")
	}

	verifiedDeleteResult, errorDeleteResult := Read(&readParams)
	if errorDeleteResult == nil {
		t.Error("Error double checking deleted password")
	}
	if verifiedDeleteResult != nil {
		t.Error("Deleted password is still present")
	}
}
