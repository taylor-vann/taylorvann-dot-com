package roles

import (
	"fmt"

	"webapi/controllers/roles/constants"
)

// SQL - container of valid SQL Rolesss
type SQL struct {
	CreateTable  string
	Create       string
	Index				 string
	Read         string
	Search			 string
	Update       string
	UpdateAccess string
	Delete       string
	Undelete		 string
}

type StatementMap = map[string]SQL

const createTableRoles = `
CREATE TABLE IF NOT EXISTS %s (
	id BIGSERIAL PRIMARY KEY,
	user_id BIGINT NOT NULL,
	organization VARCHAR(256),
	read_access BOOLEAN NOT NULL DEFAULT FALSE,
	write_access BOOLEAN NOT NULL DEFAULT FALSE,
	is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
	created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
	updated_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
	CONSTRAINT unique_role UNIQUE (user_id, organization)
);
`

const insert = `
INSERT INTO %s (
	user_id,
	organization,
	read_access,
	write_access
)
VALUES
	($1, $2, $3, $4)
RETURNING
	*;
`

const index = `
SELECT
  *
FROM
  %s
WHERE
  id BETWEEN $1 AND ($1 + $2);
`

const read = `
SELECT
	*
FROM
	%s
WHERE
	user_id = $1 AND
	organization = $2;
`

const search = `
SELECT
  *
FROM
  %s
WHERE
  user_id = $1;
`

const update = `
UPDATE
	%s
SET
	read_access = $2,
	write_access = $3,
	is_deleted = $4,
  updated_at = CURRENT_TIMESTAMP(3)
WHERE
	user_id = $1 AND
	organization = $2 AND
	TO_TIMESTAMP($7::DOUBLE PRECISION * 0.001) 
		BETWEEN updated_at AND CURRENT_TIMESTAMP(3)
RETURNING 
	*;
`

const updateAccess = `
UPDATE
	%s
SET
	read_access = $2,
	write_access = $3,
  updated_at = CURRENT_TIMESTAMP(3)
WHERE
	user_id = $1 AND
	organization = $2 AND
	TO_TIMESTAMP($7::DOUBLE PRECISION * 0.001) 
		BETWEEN updated_at AND CURRENT_TIMESTAMP(3)
RETURNING 
	*;
`

const updateAsDeleted = `
UPDATE
  %s
SET
  is_deleted = TRUE,
  updated_at = CURRENT_TIMESTAMP(3)
WHERE
	user_id = $1 AND
	organization = $2 AND
	TO_TIMESTAMP($2::DOUBLE PRECISION * 0.001) 
		BETWEEN updated_at AND CURRENT_TIMESTAMP(3)
RETURNING
	*;
`

const undelete = `
UPDATE
  %s
SET
  is_deleted = FALSE,
  updated_at = CURRENT_TIMESTAMP(3)
WHERE
	user_id = $1 AND
	organization = $2 AND
	TO_TIMESTAMP($2::DOUBLE PRECISION * 0.001) 
		BETWEEN updated_at AND CURRENT_TIMESTAMP(3)
RETURNING
	*;
`

func createRolesStatements(tableName string) SQL {
	Roles := SQL{
		CreateTable:	fmt.Sprintf(createTable, tableName),
		Create:				fmt.Sprintf(insert, tableName),
		Read:					fmt.Sprintf(read, tableName),
		Search:				fmt.Sprintf(search, tableName),
		Update:				fmt.Sprintf(update, tableName),
		UpdateAccess: fmt.Sprintf(updateAccess, tableName),
		Delete:				fmt.Sprintf(updateAsDeleted, tableName),
		Undelete:			fmt.Sprintf(undelete, tableName)
	}

	return &Roles
}

func createStatementMap() StatementMap {
	sqlMap := make(StatementMap)
	
	sqlMap[constants.Production] = createRolesStatements(constants.Users)
	sqlMap[constants.Development] = createRolesStatements(constants.UsersTest)
	sqlMap[constants.Local] = createRolesStatements(constants.UsersUnitTests)

	return sqlMap
}

var SQLStatements = createStatementMap()

const DangerouslyDropUnitTestsTable =`
DROP TABLE roles_unit_tests;
`