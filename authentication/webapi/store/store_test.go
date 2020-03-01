package store

import (
	"fmt"
	"testing"
)

const originalEmail = "test_email@email.test"
const originalPassword = "pazzwerd"
const updatedEmail = "updated_test_email@email.test"
const updatedPassword = "icanhazcheeze"
const incorrectPassword = "fizzbuzz"
const nonExistentEmail = "non_existent_test_email@email.test"

var testUser = CreateUserParams{
	Email:    originalEmail,
	Password: originalPassword,
}

var validateTestUser = ValidateUserParams{
	Email:    originalEmail,
	Password: originalPassword,
}

var readTestUser = ReadUserParams{
	Email: originalEmail,
}

var readNonExistentTestUser = ReadUserParams{
	Email: nonExistentEmail,
}

var validateIncorrectTestUser = ValidateUserParams{
	Email:    originalEmail,
	Password: incorrectPassword,
}

var updatedTestUser = UpdateEmailParams{
	CurrentEmail: originalEmail,
	UpdatedEmail: updatedEmail,
}

var updatedUserPassword = UpdatePasswordParams{
	Email:           updatedEmail,
	UpdatedPassword: updatedPassword,
}

var validateUpdatedUserPassword = ValidateUserParams{
	Email:    updatedEmail,
	Password: updatedPassword,
}

var removeTestUser = RemoveUserParams{
	Email: updatedEmail,
}

func TestCreateTables(t *testing.T) {
	_, errUsers := CreateRequiredTables()
	if errUsers != nil {
		t.Error("error creating tables from store")
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
		return
	}
	if userRow.Email != readTestUser.Email {
		t.Error("incorrect test user retrieved")
	}
}

func TestReadNonExistantUser(t *testing.T) {
	userRowNonExistent, errUser := ReadUser(&readNonExistentTestUser)
	if errUser != nil {
		fmt.Println(errUser)
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
	if correctResult == nil {
		t.Error("Correct password was not correct")
	}

	incorrectResult, errIncorrectPassword := ValidateUser(
		&validateIncorrectTestUser,
	)
	if errIncorrectPassword != nil {
		t.Error("Error verifying incorrect password")
	}
	if incorrectResult != nil {
		t.Error("Incorrect password was correct correct")
	}
}

// Update a user email
func TestUpdateEmail(t *testing.T) {
	result, errUpdateUser := UpdateEmail(&updatedTestUser)
	if errUpdateUser != nil {
		fmt.Println(errUpdateUser)
		t.Error("Error updating user")
	}

	if result == nil {
		t.Error("nil user returned")
		return
	}

	if result.Email != updatedTestUser.UpdatedEmail {
		t.Error("User did not update")
	}
}

// Update a users password
func TestUpdateUserPassword(t *testing.T) {
	result, errUpdateUserPassword := UpdatePassword(
		&updatedUserPassword,
	)

	if errUpdateUserPassword != nil {
		t.Error("Error updating user")
	}

	if result == nil {
		t.Error("nil user returned, no password to update")
		return
	}

	if result.Email != updatedTestUser.UpdatedEmail {
		t.Error("User's password did not update")
	}

	resultValidated, errValidateUserValidateUser := ValidateUser(
		&validateUpdatedUserPassword,
	)
	if errValidateUserValidateUser != nil {
		t.Error("New password cannot be validated")
	}

	if resultValidated == nil {
		t.Error("could not validate updated password")
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


func TestReviveUser(t *testing.T) {
	user, errRemoveUser := ReviveUser(
		&removeTestUser,
	)

	if errRemoveUser != nil {
		t.Error("Error removing user")
	}

	if user == nil {
		t.Error("nil result returned")
		return
	}
	if user.IsDeleted != false {
		t.Error("User is_deleted should be false")
	}

	if user.Email != removeTestUser.Email {
		t.Error("Wrong user was deleted")
	}
}