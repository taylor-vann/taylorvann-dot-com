// brian taylor vann
// taylorvann dot com

package passwords

import (
	"fmt"

	"webapi/constants"
)

// PasswordsSQL - container for valid SQL Passwords
type PasswordsSQL struct {
	CreateTable string
	Create      string
	Read        string
	Update      string
	Remove      string
}

const password = "passwords"
const passwordTest = "passwords_test"

// sCreateTablePassword - create table Passwords
const createTablePasswords = `
CREATE TABLE IF NOT EXISTS %s (
	id BIGSERIAL PRIMARY KEY,
	salt VARCHAR(1024) NOT NULL,
	hash VARCHAR(1024) NOT NULL,
	params VARCHAR(2048) NOT NULL,
  created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  updated_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3)
);
`

// InsertPassword - create a  entry
const insertPassword = `
INSERT INTO 
	%s (
  	salt,
    hash,
		params
  )
VALUES
	($1, $2, $3)
RETURNING
	*;
`

// readPassword - remove  entry
const readPassword = `
SELECT
 *
FROM
	%s
WHERE
	id = $1;
`

// UpdatePassword - update  entry
const updatePassword = `
UPDATE
	%s
SET
  salt = $2,
	hash = $3,
	params = $4,
  updated_at = CURRENT_TIMESTAMP(3)
WHERE
	id = $1 AND
	TO_TIMESTAMP($5::DOUBLE PRECISION * 0.001) 
		BETWEEN updated_at AND CURRENT_TIMESTAMP(3)
RETURNING 
	*;
`

// RemovePassword - remove  entry
const removePassword = `
DELETE FROM
	%s
WHERE
	id = $1 AND
	TO_TIMESTAMP($2::DOUBLE PRECISION * 0.001)
		BETWEEN updated_at AND CURRENT_TIMESTAMP(3)
RETURNING 
	*;
`

func createPasswords(environment string) *PasswordsSQL {
	tableName := passwordTest
	if environment == constants.Production {
		tableName = password
	}

	Passwords := PasswordsSQL{
		CreateTable: fmt.Sprintf(createTablePasswords, tableName),
		Create:      fmt.Sprintf(insertPassword, tableName),
		Read:        fmt.Sprintf(readPassword, tableName),
		Update:      fmt.Sprintf(updatePassword, tableName),
		Remove:      fmt.Sprintf(removePassword, tableName),
	}

	return &Passwords
}

// SQLStatements - interface to production SQL Passwords
var SQLStatements = createPasswords(password)
