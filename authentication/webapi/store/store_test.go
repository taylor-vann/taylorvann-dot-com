package store

import (
	"fmt"
	"testing"

	"webapi/controllers/haspassword"
	"webapi/controllers/passwords"
	"webapi/controllers/users"
)

const originalEmail = "test_email@email.test"
const originalPassword = "pazzwerd"
const updatedEmail = "updated_test_email@email.test"
const updatedPassword = "icanhazcheeze"
const incorrectPassword = "fizzbuzz"
const nonExistentEmail = "non_existent_test_email@email.test"

var testUser = CreateUserParams{
	Email:    originalEmail,
	Username: "buster",
	Password: originalPassword,
}

var validateTestUser = ValidateUserParams{
	Email:    originalEmail,
	Password: originalPassword,
}

var readTestUser = ReadUserParams{
	Email:    originalEmail,
}

var readNonExistentTestUser = ReadUserParams{
	Email:    nonExistentEmail,
}

var validateIncorrectTestUser = ValidateUserParams{
	Email:    originalEmail,
	Password: incorrectPassword,
}

var updatedTestUser = UpdateUserParams{
	CurrentEmail: originalEmail,
	UpdatedEmail: updatedEmail,
}

var updatedUserPassword = UpdateUserPasswordParams{
	Email: updatedEmail,
	UpdatedPassword: updatedPassword,
}

var validateUpdatedUserPassword = ValidateUserParams{
	Email: updatedEmail,
	Password: updatedPassword,
}

var removeTestUser = RemoveUserParams{
	Email:    updatedEmail,
}

func TestCreateTables(t *testing.T) {
	_, errUsers := users.CreateTable()
	if errUsers != nil {
		t.Error("Error creating users table")
	}

	_, errPasswords := passwords.CreateTable()
	if errPasswords != nil {
		t.Error("Error creating passwords table")
	}

	_, errHasPassword := haspassword.CreateTable()
	if errHasPassword != nil {
		t.Error("Error creating haspassword table")
	}
}

func TestCreateUser(t *testing.T) {
	userRow, err := CreateUser(&testUser)
	// return user
	if userRow == nil {
		t.Error("nil reference to user")
	}
	if err != nil {
		t.Error("Error creating user")
	}
	if userRow != nil && userRow.Email != testUser.Email {
		t.Error("correct user was not created")
	}
}

func TestReadUser(t *testing.T) {
	userRow, err := ReadUser(&readTestUser)
	if err != nil {
		t.Error("Error reading user")
	}
	if userRow == nil {
		t.Error("test user is nil")
	}
	if userRow.Email != readTestUser.Email {
		t.Error("incorrect test user retrieved")
	}

	userRowNonExistent, errUser := ReadUser(&readNonExistentTestUser)
	if errUser != nil {
		t.Error("Error reading non existent user")
	}
	if userRowNonExistent != nil {
		t.Error("non existent test user should be nil")
	}
}

func TestValidateUser(t *testing.T) {
	correctResult, errCorrectPassword := ValidateUser(&validateTestUser)
	if errCorrectPassword != nil {
		t.Error("Error verifying correct password")
	}
	if correctResult == false {
		t.Error("Correct password was not correct")
	}

	incorrectResult, errIncorrectPassword := ValidateUser(
		&validateIncorrectTestUser,
	)
	if errIncorrectPassword != nil {
		t.Error("Error verifying incorrect password")
	}
	if incorrectResult == true {
		t.Error("Incorrect password was correct correct")
	}
}

// Update a user email
func TestUpdatedUser(t *testing.T) {
	result, errUpdateUser := UpdateUser(&updatedTestUser)
	if errUpdateUser != nil {
		fmt.Println(errUpdateUser)
		t.Error("Error updating user")
	}

	if result != nil && result.Email != updatedTestUser.UpdatedEmail {
		t.Error("User did not update")
	}

	if result == nil {
		t.Error("nil user returned")
	}
}

// Update a users password
func TestUpdateUserPassword(t *testing.T) {
	result, errUpdateUserPassword := UpdateUserPassword(
		&updatedUserPassword,
	)

	if errUpdateUserPassword != nil {
		fmt.Println(errUpdateUserPassword)
		t.Error("Error updating user")
	}

	if result != nil && result.Email != updatedTestUser.UpdatedEmail {
		t.Error("User's password did not update")
	}

	if result == nil {
		t.Error("nil user returned, no password to update")
	}

	isValidated, errValidateUserValidateUser := ValidateUser(
		&validateUpdatedUserPassword,
	)
	if errValidateUserValidateUser != nil {
		t.Error("New password cannot be validated")
	}
	if isValidated == false {
		t.Error("Could not find a user to validate")
	}
}

// Remove a user

func TestRemoveUser(t *testing.T) {
	user, errRemoveUser := RemoveUser(
		&removeTestUser,
	)

	if user == nil {
		t.Error("nil result returned")
	}

	if errRemoveUser != nil {
		t.Error("Error removing user")
	}

	if user != nil && user.IsDeleted == false {
		t.Error("User is_deleted should be true")
	}

	if user != nil && user.Email != removeTestUser.Email {
		t.Error("Wrong user was deleted")
	}
}

