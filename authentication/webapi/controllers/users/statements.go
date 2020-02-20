// Package users - statements required of users table
package users

import (
	"fmt"

	"webapi/controllers/users/constants"
)

// SQL - container of valid sql strings
type SQL struct {
	CreateTable string
	Create      string
	Read        string
	Update      string
	Remove      string
}

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

const readUser = `
SELECT
  *
FROM
  %s
WHERE
	email = $1;
`

const updateUser = `
UPDATE
	%s
SET
	email = $2,
	salt = $3,
	hash = $4,
	params = $5,
  updated_at = CURRENT_TIMESTAMP(3)
WHERE
	email = $1 AND
	TO_TIMESTAMP($6::DOUBLE PRECISION * 0.001) 
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
	TO_TIMESTAMP($2::DOUBLE PRECISION * 0.001) 
		BETWEEN updated_at AND CURRENT_TIMESTAMP(3)
RETURNING
	*;
`

func createUsersStatements() *SQL {
	Users := SQL{
		CreateTable: fmt.Sprintf(createTableUsers, constants.Tables.Users),
		Create:      fmt.Sprintf(insertUser, constants.Tables.Users),
		Read:        fmt.Sprintf(readUser, constants.Tables.Users),
		Update:      fmt.Sprintf(updateUser, constants.Tables.Users),
		Remove:      fmt.Sprintf(updateAsDeletedUser, constants.Tables.Users),
	}

	return &Users
}

// SQLStatements - interface to production SQL Userss
var SQLStatements = createUsersStatements()
