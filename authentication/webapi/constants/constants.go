package constants

import (
	"os"
)

// TableList -
type TableList struct {
	Users				string
	HasPassword string
	Passwords		string
}

// Stage -
const Stage = "STAGE"

// Environment
var Environment = os.Getenv(Stage)

// Production -
const Production = "PRODUCTION"

// Tables
// Users -
const Users = "users"

// UsersTest -
const UsersTest = "users_test"

// Passwords -
const Passwords = "passwords"

// PasswordsTest -
const PasswordsTest = "passwords_test"

// HasPassword -
const HasPassword = "haspassword"

// HasPasswordTest -
const HasPasswordTest = "haspassword_test"

// Tables Based on Environment
// UsersTable -
var UsersTable = UsersTest
// PasswordsTable -
var PasswordsTable = PasswordsTest
// HasPasswordTable -
var HasPasswordTable = HasPasswordTest

// Tables
var Tables = SetEnvironmentConstants()

// SetEnvironmentConstants -
func SetEnvironmentConstants() *TableList {
	var tables TableList
	tables.Users = UsersTest
	tables.HasPassword = HasPasswordTest
	tables.Passwords = PasswordsTest

	if Environment == Production {
		tables.Users = Users
		tables.Passwords = Passwords
		tables.HasPassword = HasPassword
	}

	return &tables
}
