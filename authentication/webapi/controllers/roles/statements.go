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
	user_id BIGINT UNIQUE NOT NULL,
	role VARCHAR(128),
	created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
	updated_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
	CONSTRAINT unique_role UNIQUE (user_id, role)
);
`

const insertRole = `
INSERT INTO 
	%s (
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
	user_id = $1;
`

const updateRole = `
UPDATE
	%s
SET
  role = $2,
  updated_at = CURRENT_TIMESTAMP(3)
WHERE
	user_id = $1 AND
	TO_TIMESTAMP($3::DOUBLE PRECISION * 0.001) 
		BETWEEN updated_at AND CURRENT_TIMESTAMP(3)
RETURNING 
	*;
`

const removeRole = `
DELETE FROM
	%s
WHERE
	user_id = $1 AND
	TO_TIMESTAMP($2::DOUBLE PRECISION * 0.001) 
		BETWEEN updated_at AND CURRENT_TIMESTAMP(3)
RETURNING 
	*;
`

func createRoles() *SQL {
	Roles := SQL{
		CreateTable: fmt.Sprintf(createTableRoles, constants.Tables.Roles),
		Create:      fmt.Sprintf(insertRole, constants.Tables.Roles),
		Read:        fmt.Sprintf(readRole, constants.Tables.Roles),
		Update:      fmt.Sprintf(updateRole, constants.Tables.Roles),
		Remove:      fmt.Sprintf(removeRole, constants.Tables.Roles),
	}

	return &Roles
}

// SQLStatements - interface to production SQL Roless
var SQLStatements = createRoles()
