// brian taylor vann
// taylorvann dot com

// store n
package store

import (
	// "encoding/json"
	// "errors"
	"time"

	rolesController "webapi/store/roles/controller"
	usersController "webapi/store/users/controller"
	// "webapi/interfaces/passwordx"
	// "webapi/interfaces/storex"
)

// MilliSeconds -
type MilliSeconds = int64

// CreateUserParams -
type CreateUserParams = usersController.CreateParams

// ReadUserParams -
type ReadUserParams = usersController.ReadParams

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
type UserRow = usersController.Row

// getNowAsMS -
func getNowAsMS() MilliSeconds {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// CreateRequiredTables -
func CreateRequiredTables() (bool, error) {
	_, errDevelopment := usersController.CreateTable(&usersController.CreateTableParams{
		Environment: "DEVELOPMENT",
	})
	if errDevelopment != nil {
		return false, errDevelopment
	}

	_, errProduction := usersController.CreateTable(&usersController.CreateTableParams{
		Environment: "PRODUCTION",
	})
	if errProduction != nil {
		return false, errProduction
	}

	_, errRolesDevelopment := rolesController.CreateTable(&rolesController.CreateTableParams{
		Environment: "DEVELOPMENT",
	})
	if errRolesDevelopment != nil {
		return false, errRolesDevelopment
	}

	_, errRolesProduction := rolesController.CreateTable(&rolesController.CreateTableParams{
		Environment: "PRODUCTION",
	})
	if errRolesProduction != nil {
		return false, errRolesProduction
	}

	return true, nil
}
