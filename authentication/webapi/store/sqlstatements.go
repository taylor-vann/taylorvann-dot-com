package store

import (
	"fmt"
	"os"
	"webapi/constants"
)

// DirectStoreSQL -
type DirectStoreSQL struct {
	InsertUserAndPassword string
}

const insertUserAndPassword = `
WITH 
	user_result AS (
		INSERT INTO %s (
			email
		)
		VALUES (
			$1
		)
		RETURNING
			*
	),
	password_result AS (
			INSERT INTO %s (
				salt,
				hash,
				params
			)
			VALUES
				($2, $3, $4)
			RETURNING
				*
	),
	haspassword_result AS (
		INSERT INTO %s (
			user_id,
			password_id
		)
		VALUES
			(
				(SELECT id FROM user_result LIMIT 1),
				(SELECT id FROM password_result LIMIT 1)
			)
		RETURNING
			*
	)
	SELECT * 
	FROM user_result 
	LIMIT 1;
`

func formatInsertUserAndPassword(environment string) string {
	usersTable := constants.UsersTest
	passwordsTable := constants.PasswordsTest
	haspasswordTable := constants.HasPasswordTest

	if environment == constants.Production {
		usersTable = constants.Users
		passwordsTable = constants.Passwords
		haspasswordTable = constants.HasPassword
	}

	return fmt.Sprintf(
		insertUserAndPassword,
		usersTable,
		passwordsTable,
		haspasswordTable,
	)
}

var environment = os.Getenv(constants.Stage)
var formattedInsertUserAndPassword = formatInsertUserAndPassword(environment)

// SQLStatements -
var SQLStatements = DirectStoreSQL{
	InsertUserAndPassword: formattedInsertUserAndPassword,
}
