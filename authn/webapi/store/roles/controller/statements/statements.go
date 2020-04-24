// brian taylor vann
// taylorvann dot com

package statements

import (
	"fmt"

	"webapi/store/roles/controller/constants"
)

type SQL struct {
	CreateTable  string
	Create       string
	Read         string
	Search			 string
	Index				 string
	Update       string
	UpdateAccess string
	Delete       string
	Undelete		 string
}

type StatementMap = map[string]SQL

const createTable = `
CREATE TABLE IF NOT EXISTS %s (
	id BIGSERIAL PRIMARY KEY,
	user_id BIGINT NOT NULL,
	organization VARCHAR(256) NOT NULL,
	read_access BOOLEAN NOT NULL DEFAULT FALSE,
	write_access BOOLEAN NOT NULL DEFAULT FALSE,
	is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
	created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
	updated_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
	CONSTRAINT uc_user_and_organization UNIQUE (user_id, organization)
);
`

const create = `
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
	user_id = $1
OFFSET
	$2
LIMIT
	($2 + $3)
`

const update = `
UPDATE
	%s
SET
	read_access = $3,
	write_access = $4,
	is_deleted = $5,
  updated_at = CURRENT_TIMESTAMP(3)
WHERE
	user_id = $1 AND
	organization = $2 AND
	TO_TIMESTAMP($6::DOUBLE PRECISION * 0.001) 
		BETWEEN updated_at AND CURRENT_TIMESTAMP(3)
RETURNING 
	*;
`

const updateAccess = `
UPDATE
	%s
SET
	read_access = $3,
	write_access = $4,
  updated_at = CURRENT_TIMESTAMP(3)
WHERE
	user_id = $1 AND
	organization = $2 AND
	TO_TIMESTAMP($5::DOUBLE PRECISION * 0.001) 
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
	TO_TIMESTAMP($3::DOUBLE PRECISION * 0.001) 
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
	TO_TIMESTAMP($3::DOUBLE PRECISION * 0.001) 
		BETWEEN updated_at AND CURRENT_TIMESTAMP(3)
RETURNING
	*;
`

func createRolesStatements(tableName string) SQL {
	statements := SQL{
		CreateTable:	fmt.Sprintf(createTable, tableName),
		Create:				fmt.Sprintf(create, tableName),
		Read:					fmt.Sprintf(read, tableName),
		Search:				fmt.Sprintf(search, tableName),
		Index:				fmt.Sprintf(index, tableName),
		Update:				fmt.Sprintf(update, tableName),
		UpdateAccess: fmt.Sprintf(updateAccess, tableName),
		Delete:				fmt.Sprintf(updateAsDeleted, tableName),
		Undelete:			fmt.Sprintf(undelete, tableName),
	}

	return statements
}

func createStatementMap() StatementMap {
	sqlMap := make(StatementMap)
	
	sqlMap[constants.Production] = createRolesStatements(constants.Roles)
	sqlMap[constants.Development] = createRolesStatements(constants.RolesTest)
	sqlMap[constants.Local] = createRolesStatements(constants.RolesUnitTests)

	return sqlMap
}

var SqlMap = createStatementMap()

const DangerouslyDropUnitTestsTable =`
DROP TABLE roles_unit_tests;
`