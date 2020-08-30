// brian taylor vann
// briantaylorvann dot com

package store

import (
	"os"

	rolesController "webapi/store/roles/controller"
	usersController "webapi/store/users/controller"
)

const (
	local       = "LOCAL"
	development = "DEVELOPMENT"
	production  = "PRODUCTION"
)

var Environment = os.Getenv("STAGE")

func CreateLocalTables() (bool, error) {
	_, errLocal := usersController.CreateTable(&usersController.CreateTableParams{
		Environment: local,
	})
	if errLocal != nil {
		return false, errLocal
	}

	_, errRolesLocal := rolesController.CreateTable(&rolesController.CreateTableParams{
		Environment: local,
	})
	if errRolesLocal != nil {
		return false, errRolesLocal
	}

	return true, nil
}

func CreateRequiredTables() (bool, error) {
	_, errDevelopment := usersController.CreateTable(&usersController.CreateTableParams{
		Environment: development,
	})
	if errDevelopment != nil {
		return false, errDevelopment
	}

	_, errProduction := usersController.CreateTable(&usersController.CreateTableParams{
		Environment: production,
	})
	if errProduction != nil {
		return false, errProduction
	}

	_, errRolesDevelopment := rolesController.CreateTable(&rolesController.CreateTableParams{
		Environment: development,
	})
	if errRolesDevelopment != nil {
		return false, errRolesDevelopment
	}

	_, errRolesProduction := rolesController.CreateTable(&rolesController.CreateTableParams{
		Environment: production,
	})
	if errRolesProduction != nil {
		return false, errRolesProduction
	}

	if Environment == local {
		return CreateLocalTables()
	}

	return true, nil
}
