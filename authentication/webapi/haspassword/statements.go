// brian taylor vann
// taylorvann dot com

// postgresql statements required of haspassword table

package haspassword

// CreatePasswordsTableStatement - Create table Statement
const CreatePasswordsTableStatement = `
CREATE TABLE IF NOT EXISTS has_password (
	id BIGSERIAL PRIMARY KEY,
	user_id BIGINT UNIQUE NOT NULL,
	password_id BIGINT UNIQUE NOT NULL,
  created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3)
)
`

// InsertHasPasswordStatement - Create a password
const InsertHasPasswordStatement = `
INSERT INTO 
	has_password (
  	user_id,
    password_id
  )
VALUES
	($1, $2)
RETURNING
	*;
`

// DeleteHasPasswordStatement - Update personal password salt and hash
const DeleteHasPasswordStatement = `
DELETE FROM
	has_password
WHERE
	id = $1
RETURNING 
	*;
`

// Create
// Delete
