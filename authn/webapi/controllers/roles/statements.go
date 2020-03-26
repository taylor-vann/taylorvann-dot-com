package roles

import (
	"fmt"

	"webapi/controllers/roles/constants"
)

// SQL - container of valid SQL Rolesss
type SQL struct {
	CreateTable string
	Create      string
	Read        string
	Update      string
	Remove      string
}

// CreateTableRoles - Create table Roles
const createTableRoles = `
CREATE TABLE IF NOT EXISTS %s (
	id BIGSERIAL PRIMARY KEY,
	user_id BIGINT NOT NULL,
	role VARCHAR(128),
	created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
	CONSTRAINT unique_role UNIQUE (user_id, role)
);
`

const insertRole = `
INSERT INTO %s (
	user_id,
	role
)
VALUES
	($1, $2)
RETURNING
	*;
`

const readRole = `
SELECT
	*
FROM
	%s
WHERE
	user_id = $1 AND
	role = $2;
`

const removeRole = `
DELETE FROM
	%s
WHERE
	user_id = $1 AND
	role = $2
RETURNING 
	*;
`

func createRoles() *SQL {
	Roles := SQL{
		CreateTable: fmt.Sprintf(createTableRoles, constants.Tables.Roles),
		Create:      fmt.Sprintf(insertRole, constants.Tables.Roles),
		Read:        fmt.Sprintf(readRole, constants.Tables.Roles),
		Remove:      fmt.Sprintf(removeRole, constants.Tables.Roles),
	}

	return &Roles
}

// SQLStatements - interface to production SQL Roless
var SQLStatements = createRoles()
