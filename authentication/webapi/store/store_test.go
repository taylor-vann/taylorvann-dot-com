package store

import (
	"fmt"
	"testing"
	"webapi/controllers/haspassword"
	"webapi/controllers/passwords"
	"webapi/controllers/users"
)

var testUser = CreateUserParams{
	Email:    "test_email@email.test",
	Username: "buster",
	Password: "pazzwerd",
}

var validateTestUser = ValidateUserParams{
	Email:    "test_email@email.test",
	Password: "pazzwerd",
}

var validateIncorrectTestUser = ValidateUserParams{
	Email:    "test_email@email.test",
	Password: "fizzbuzz",
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
	if err != nil {
		fmt.Println(err)
		t.Error("Error creating user")
	}
	fmt.Print(userRow)
	t.Error("Fail because it's new")
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

// Update a user

// Update a users password

// Update a users email <- public? no you wouldn't publish emails

// Update a users username?

// Remove a user

// Validate a user

