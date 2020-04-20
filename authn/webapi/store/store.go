// brian taylor vann
// taylorvann dot com

// store n
package store

import (
	// "encoding/json"
	// "errors"
	"time"

	"webapi/controllers/roles"
	"webapi/controllers/users"
	// "webapi/interfaces/passwordx"
	// "webapi/interfaces/storex"
)

// MilliSeconds -
type MilliSeconds = int64

// CreateUserParams -
type CreateUserParams = users.CreateParams

// ReadUserParams -
type ReadUserParams = users.ReadParams

// UpdateEmailParams -
type UpdateEmailParams struct {
	CurrentEmail string
	UpdatedEmail string
}

// ValidateUserParams -
type ValidateUserParams struct {
	Email    string
	Password string
}

// RemoveUserParams -
type RemoveUserParams = ReadUserParams

// ReviveUserParams -
type ReviveUserParams = ReadUserParams

// UpdatePasswordParams -
type UpdatePasswordParams struct {
	Email           string
	UpdatedPassword string
}

// UserRow -
type UserRow = users.Row

// getNowAsMS -
func getNowAsMS() MilliSeconds {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// CreateRequiredTables -
func CreateRequiredTables() (bool, error) {
	_, errDevelopment := users.CreateTable(&users.CreateTableParams{
		Environment: "DEVELOPMENT",
	})
	if errDevelopment != nil {
		return false, errDevelopment
	}

	_, errProduction := users.CreateTable(&users.CreateTableParams{
		Environment: "PRODUCTION",
	})
	if errProduction != nil {
		return false, errProduction
	}

	_, errRolesDevelopment := roles.CreateTable(&roles.CreateTableParams{
		Environment: "DEVELOPMENT",
	})
	if errRolesDevelopment != nil {
		return false, errRolesDevelopment
	}

	_, errRolesProduction := roles.CreateTable(&roles.CreateTableParams{
		Environment: "PRODUCTION",
	})
	if errRolesProduction != nil {
		return false, errRolesProduction
	}

	return true, nil
}
