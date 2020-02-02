// Package statements - statements required of users table
package statements

import "fmt"

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

func createUsers(tableName string) *UsersSQL {
	Users := UsersSQL{
		CreateTable: fmt.Sprintf(createTableUsers, tableName),
		Create:      fmt.Sprintf(insertUser, tableName),
		Read:        fmt.Sprintf(readUser, tableName),
		Update:      fmt.Sprintf(updateUser, tableName),
		Remove:      fmt.Sprintf(updateAsDeletedUser, tableName),
	}

	return &Users
}

// Users - interface to production SQL Userss
var Users = createUsers(users)

// UsersTest - interface to development SQL Userss
var UsersTest = createUsers(usersTest)
