package haspassword

import (
	"fmt"

	"webapi/constants"
)

// HasPasswordsSQL - container of valid SQL HasPasswordss
type HasPasswordsSQL struct {
	CreateTable string
	Create      string
	Read        string
	Update      string
	Remove      string
}

// CreateTableHasPasswords - Create table HasPasswords
const createTableHasPassword = `
CREATE TABLE IF NOT EXISTS %s (
	id BIGSERIAL PRIMARY KEY,
	user_id BIGINT UNIQUE NOT NULL,
	password_id BIGINT UNIQUE NOT NULL,
	created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
	updated_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3)
);
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

const updateHasPassword = `
UPDATE
	%s
SET
  password_id = $2,
  updated_at = CURRENT_TIMESTAMP(3)
WHERE
	user_id = $1 AND
	TO_TIMESTAMP($3::DOUBLE PRECISION * 0.001) 
		BETWEEN updated_at AND CURRENT_TIMESTAMP(3)
RETURNING 
	*;
`

const removeHasPassword = `
DELETE FROM
	%s
WHERE
	user_id = $1 AND
	TO_TIMESTAMP($2::DOUBLE PRECISION * 0.001) 
		BETWEEN updated_at AND CURRENT_TIMESTAMP(3)
RETURNING 
	*;
`

func createHasPassword() *HasPasswordsSQL {
	HasPasswords := HasPasswordsSQL{
		CreateTable: fmt.Sprintf(createTableHasPassword, constants.Tables.HasPassword),
		Create:      fmt.Sprintf(insertHasPassword, constants.Tables.HasPassword),
		Read:        fmt.Sprintf(readHasPassword, constants.Tables.HasPassword),
		Update:      fmt.Sprintf(updateHasPassword, constants.Tables.HasPassword),
		Remove:      fmt.Sprintf(removeHasPassword, constants.Tables.HasPassword),
	}

	return &HasPasswords
}

// SQLStatements - interface to production SQL HasPasswords
var SQLStatements = createHasPassword()
