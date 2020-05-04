// brian taylor vann
// taylorvann dot com

package store

import (
	"time"

	rolesController "webapi/store/roles/controller"
	usersController "webapi/store/users/controller"
)

type MilliSeconds = int64
type CreateUserParams = usersController.CreateParams
type ReadUserParams = usersController.ReadParams
type UpdateEmailParams struct {
	CurrentEmail string
	UpdatedEmail string
}

type ValidateUserParams struct {
	Email    string
	Password string
}

type RemoveUserParams = ReadUserParams
type ReviveUserParams = ReadUserParams
type UpdatePasswordParams struct {
	Email           string
	UpdatedPassword string
}

type UserRow = usersController.Row

func getNowAsMS() MilliSeconds {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

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
