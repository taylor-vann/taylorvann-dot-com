package passwords

import (
	"webapi/passwords"
)

// CreatePasswordsTableStatement - Create table Statement
const CreatePasswordsTableStatement = `
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	salt VARCHAR(1024) NOT NULL,
	hash VARCHAR(1024) NOT NULL,
	params VARCHAR(2048) NOT NULL,
  created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  updated_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3)
)
`

// InsertPasswordStatement - Create a password
const InsertPasswordStatement = `
INSERT INTO passwords (
  	salt,
    hash,
		params,
  )
VALUES
	($1, $2, $3)
RETURNING
	*;
`

// UpdatePasswordStatement - Update personal password salt and hash
const UpdatePasswordStatement = `
UPDATE
	passwords
SET
  salt = $2,
	hash = $3,
	params = $4,
  updated_at = CURRENT_TIMESTAMP(3)
WHERE
	id = $1
RETURNING 
	*;
`