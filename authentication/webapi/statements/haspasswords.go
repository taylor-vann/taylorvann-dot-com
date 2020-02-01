// Package statements - HasPasswordss required of has_password table
package statements

import "fmt"

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

func createHasPasswords(tableName string) *HasPasswordsSQL {
	HasPasswords := HasPasswordsSQL{
		CreateTable: fmt.Sprintf(createTableHasPasswords, tableName),
		Create:      fmt.Sprintf(insertHasPassword, tableName),
		Read:        fmt.Sprintf(readHasPassword, tableName),
		Remove:      fmt.Sprintf(removeHasPassword, tableName),
	}

	return &HasPasswords
}

// HasPasswords - interface to production SQL HasPasswordss
var HasPasswords = createHasPasswords(hasPassword)

// HasPasswordsTest - interface to development SQL HasPasswordss
var HasPasswordsTest = createHasPasswords(hasPasswordTest)
