// Package users - statements required of users table
package users

import (
	"fmt"

	"webapi/constants"
)

// UsersSQL - container of valid sql strings
type UsersSQL struct {
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
  is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
	updated_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3)
);
`

const insertUser = `
INSERT INTO %s (
	email
)
VALUES 
  ($1)
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
  updated_at = CURRENT_TIMESTAMP(3)
WHERE
	email = $1 AND
	TO_TIMESTAMP($3::DOUBLE PRECISION * 0.001) 
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

func createUsersStatements() *UsersSQL {
	Users := UsersSQL{
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
