// brian taylor vann
// briantaylorvann dot com

package store

import (
	"os"

	rolesController "webapi/store/roles/controller"
	usersController "webapi/store/users/controller"
)

var Environment = os.Getenv ("STAGE")

func CreateLocalTables() (bool, error) {
	_, errLocal := usersController.CreateTable(&usersController.CreateTableParams{
		Environment: "LOCAL",
	})
	if errLocal != nil {
		return false, errLocal
	}

	_, errRolesLocal := rolesController.CreateTable(&rolesController.CreateTableParams{
		Environment: "LOCAL",
	})
	if errRolesLocal != nil {
		return false, errRolesLocal
	}

	return true, nil
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

	if Environment == "LOCAL" {
		return CreateLocalTables()
	}
	
	return true, nil
}
