package passwords

import (
	"webapi/interfaces/passwordx"
)

// PasswordHashParams - Expected hash structure
type PasswordHashParams = passwordx.HashParams

// PasswordStructure - Expected PostgreSQL structure
type PasswordStructure struct {
	userID    int
	salt      string
	hash      string
	params    PasswordHashParams
	createdAt string
	updatedAt string
}

// Controller for Passwords and Postgresql

// We want to create a password and link it to a user
// we want to update a password
// we want to update password and params