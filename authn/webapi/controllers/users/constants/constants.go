package constants

import (
	"webapi/constants"
)

// TableNames -
type TableNames struct {
	Users string
}

const (
	users     			= "users"
	usersTest 			= "users_test"
	usersUnitTests  = "users_unit_tests"
)

// Tables -
var Tables = getUsersTablesConstants()
// getUsersTablesConstants -
func getUsersTablesConstants() *TableNames {
	if constants.Environment == constants.Production {
		return &TableNames{
			Users: users,
		}
	}

	if constants.Environment == constants.Development {
		return &TableNames{
			Users: usersTest,
		}
	}

	return &TableNames{
		Users: usersUnitTests,
	}
}
