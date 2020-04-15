package statements

import (
	"fmt"

	"webapi/controllers/users/constants"
)

type SQL struct {
	CreateTable 	 string
	Create      	 string
	Read        	 string
	Search				 string
	Update      	 string
	UpdateEmail 	 string
	UpdatePassword string
	Remove      	 string
	Revive  			 string
}

type StatementMap = map[string]SQL

const createTableUsers = `
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

const insertUser = `
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

const updateAsDeletedUser = `
UPDATE
  %s
SET
  is_deleted = TRUE,
  updated_at = CURRENT_TIMESTAMP(3)
WHERE
	email = $1 AND
	is_deleted = FALSE AND
	TO_TIMESTAMP($2::DOUBLE PRECISION * 0.001) 
		BETWEEN updated_at AND CURRENT_TIMESTAMP(3)
RETURNING
	*;
`

const updateAsRevivedUser = `
UPDATE
  %s
SET
  is_deleted = FALSE,
  updated_at = CURRENT_TIMESTAMP(3)
WHERE
	email = $1 AND
	is_deleted = TRUE AND
	TO_TIMESTAMP($2::DOUBLE PRECISION * 0.001) 
		BETWEEN updated_at AND CURRENT_TIMESTAMP(3)
RETURNING
	*;
`

func createUsersStatements(tableName string) SQL {
	userStatements := SQL{
		CreateTable: 		fmt.Sprintf(createTableUsers, tableName),
		Create:      		fmt.Sprintf(insertUser, tableName),
		Read:        		fmt.Sprintf(read, tableName),
		Search:			 		fmt.Sprintf(search, tableName),
		Update:      		fmt.Sprintf(update, tableName),
		UpdateEmail:    fmt.Sprintf(updateEmail, tableName),
		UpdatePassword: fmt.Sprintf(updatePassword, tableName),
		Remove:      		fmt.Sprintf(updateAsDeletedUser, tableName),
		Revive:      		fmt.Sprintf(updateAsRevivedUser, tableName),
	}

	return userStatements
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