package constants

import (
	"webapi/constants"
)

// TableNames -
type TableNames struct {
	Roles string
}

const (
	roles     = "roles"
	rolesTest = "roles_test"
)

// Tables -
var Tables = getRolesTablesConstants()

// getrolesTablesConstants -
func getRolesTablesConstants() *TableNames {
	if constants.Environment == constants.Production {
		return &TableNames{
			Roles: roles,
		}
	}

	return &TableNames{
		Roles: rolesTest,
	}
}
