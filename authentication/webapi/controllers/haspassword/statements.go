// Package statements - HasPasswordss required of has_password table
package statements

import (
	"fmt"
	"os"

	"webapi/constants"
)

// HasPasswordsSQL - container of valid SQL HasPasswordss
type HasPasswordsSQL struct {
	CreateTable string
	Create      string
	Read        string
	Remove      string
}

const hasPassword = "has_password"
const hasPasswordTest = "has_password_test"

// CreateTableHasPasswords - Create table HasPasswords
const createTableHasPasswords = `
CREATE TABLE IF NOT EXISTS %s (
	id BIGSERIAL PRIMARY KEY,
	user_id BIGINT UNIQUE NOT NULL,
	password_id BIGINT UNIQUE NOT NULL,
  created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3)
)
`

const insertHasPassword = `
INSERT INTO 
	%s (
  	user_id,
    password_id
  )
VALUES
	($1, $2)
RETURNING
	*;
`

const readHasPassword = `
SELECT
	*
FROM
	%s
WHERE
	user_id = $1;
`

const removeHasPassword = `
DELETE FROM
	%s
WHERE
	user_id = $1
RETURNING 
	*;
`

func createHasPasswords(environment string) *HasPasswordsSQL {
	tableName := hasPasswordTest
	if environment == constants.Production {
		tableName = hasPassword
	}

	HasPasswords := HasPasswordsSQL{
		CreateTable: fmt.Sprintf(createTableHasPasswords, tableName),
		Create:      fmt.Sprintf(insertHasPassword, tableName),
		Read:        fmt.Sprintf(readHasPassword, tableName),
		Remove:      fmt.Sprintf(removeHasPassword, tableName),
	}

	return &HasPasswords
}

var envionrment = os.Getenv(constants.Stage)

// Statements - interface to production SQL HasPasswords
var SQLStatements = createHasPasswords(envionrment)
