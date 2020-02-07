// Package users - statements required of users table
package users

import (
	"fmt"
	"os"

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

const users = "users"
const usersTest = "users_test"

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
INSERT INTO %s
  (email)
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
	email = $1
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
	email = $1
RETURNING
	*;
`

func createUsersStatements(environment string) *UsersSQL {
	tableName := usersTest
	if environment == constants.Production {
		tableName = users
	}

	Users := UsersSQL{
		CreateTable: fmt.Sprintf(createTableUsers, tableName),
		Create:      fmt.Sprintf(insertUser, tableName),
		Read:        fmt.Sprintf(readUser, tableName),
		Update:      fmt.Sprintf(updateUser, tableName),
		Remove:      fmt.Sprintf(updateAsDeletedUser, tableName),
	}

	return &Users
}

var envionrment = os.Getenv(constants.Stage)

// SQLStatements - interface to production SQL Userss
var SQLStatements = createUsersStatements(envionrment)
