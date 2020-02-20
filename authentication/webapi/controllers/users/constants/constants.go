package constants

import (
	"webapi/constants"
)

// TableNames -
type TableNames struct {
	Users string
}

const (
	users     = "users"
	usersTest = "users_test"
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

	return &TableNames{
		Users: usersTest,
	}
}
