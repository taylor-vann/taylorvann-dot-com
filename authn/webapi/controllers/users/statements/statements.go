package statements

import (
	"fmt"

	"webapi/controllers/users/constants"
)

type SQL struct {
	CreateTable 	 string
	Create      	 string
	Index					 string
	Read        	 string
	Search				 string
	Update      	 string
	UpdateEmail 	 string
	UpdatePassword string
	Delete      	 string
	Undelete  		 string
}

type StatementMap = map[string]SQL

const createTable = `
CREATE TABLE IF NOT EXISTS %s (
  id BIGSERIAL PRIMARY KEY,
	email VARCHAR(512) UNIQUE NOT NULL,
	salt VARCHAR(256) NOT NULL,
	hash VARCHAR(512) NOT NULL,
	params VARCHAR(1024) NOT NULL,
	is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
	updated_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3)
);
`

const create = `
INSERT INTO %s (
	email,
	salt,
	hash,
	params
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
	email = $1;
`

const search = `
SELECT
  *
FROM
  %s
WHERE
  POSITION($1 in email) > 0;
`

const update = `
UPDATE
	%s
SET
	email = $2,
	is_deleted = $3,
	salt = $4,
	hash = $5,
	params = $6,
  updated_at = CURRENT_TIMESTAMP(3)
WHERE
	email = $1 AND
	TO_TIMESTAMP($7::DOUBLE PRECISION * 0.001) 
		BETWEEN updated_at AND CURRENT_TIMESTAMP(3)
RETURNING
	*;
`

const updateEmail = `
UPDATE
	%s
SET
	email = $2,
  updated_at = CURRENT_TIMESTAMP(3)
WHERE
	email = $1 AND
	TO_TIMESTAMP($3::DOUBLE PRECISION * 0.001) 
		BETWEEN updated_at AND CURRENT_TIMESTAMP(3)
RETURNING
	*;
`

const updatePassword = `
UPDATE
	%s
SET
	salt = $2,
	hash = $3,
	params = $4,
	updated_at = CURRENT_TIMESTAMP(3)
WHERE
	email = $1 AND
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
	email = $1 AND
	TO_TIMESTAMP($2::DOUBLE PRECISION * 0.001) 
		BETWEEN updated_at AND CURRENT_TIMESTAMP(3)
RETURNING
	*;
`

const updateAsUndeleted = `
UPDATE
  %s
SET
  is_deleted = FALSE,
  updated_at = CURRENT_TIMESTAMP(3)
WHERE
	email = $1 AND
	TO_TIMESTAMP($2::DOUBLE PRECISION * 0.001) 
		BETWEEN updated_at AND CURRENT_TIMESTAMP(3)
RETURNING
	*;
`

func createUsersStatements(tableName string) SQL {
	statements := SQL{
		CreateTable: 		fmt.Sprintf(createTable, tableName),
		Create:      		fmt.Sprintf(create, tableName),
		Index:					fmt.Sprintf(index, tableName),
		Read:        		fmt.Sprintf(read, tableName),
		Search:			 		fmt.Sprintf(search, tableName),
		Update:      		fmt.Sprintf(update, tableName),
		UpdateEmail:    fmt.Sprintf(updateEmail, tableName),
		UpdatePassword: fmt.Sprintf(updatePassword, tableName),
		Delete:      		fmt.Sprintf(updateAsDeleted, tableName),
		Undelete:      	fmt.Sprintf(updateAsUndeleted, tableName),
	}

	return statements
}

func createStatementMap() StatementMap {
	sqlMap := make(StatementMap)
	
	sqlMap[constants.Production] = createUsersStatements(constants.Users)
	sqlMap[constants.Development] = createUsersStatements(constants.UsersTest)
	sqlMap[constants.Local] = createUsersStatements(constants.UsersUnitTests)

	return sqlMap
}

var SqlMap = createStatementMap()

const DangerouslyDropUnitTestsTable =`
DROP TABLE users_unit_tests;
`