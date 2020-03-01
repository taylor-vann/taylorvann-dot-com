package store

import (
	"fmt"
	"webapi/controllers/users/constants"
)

// SQL -
type SQL struct {
	UpdateEmail			string
	UpdatePassword	string
	RemoveUser			string
	ReviveUser			string
}

const updateEmail = `
	UPDATE
		%s
	SET
		email = $2
	WHERE
		email = $1 AND
		is_deleted = FALSE AND
		TO_TIMESTAMP($3::DOUBLE PRECISION * 0.001) 
			BETWEEN updated_at AND CURRENT_TIMESTAMP(3)
	RETURNING
		*;
`

const updatePassword = `
	UPDATE
		%s
	SET
		salt = $2,
		hash = $3,
		params = $4
	WHERE
		email = $1 AND
		is_deleted = FALSE AND
		TO_TIMESTAMP($5::DOUBLE PRECISION * 0.001) 
			BETWEEN updated_at AND CURRENT_TIMESTAMP(3)
	RETURNING
		*;
`

const removeUser = `
	UPDATE
		%s
	SET
		is_deleted = TRUE
	WHERE
		email = $1 AND
		TO_TIMESTAMP($2::DOUBLE PRECISION * 0.001) 
			BETWEEN updated_at AND CURRENT_TIMESTAMP(3)
	RETURNING
		*;
`

const reviveUser = `
	UPDATE
		%s
	SET
		is_deleted = FALSE
	WHERE
		email = $1 AND
		TO_TIMESTAMP($2::DOUBLE PRECISION * 0.001) 
			BETWEEN updated_at AND CURRENT_TIMESTAMP(3)
	RETURNING
		*;
`

// SQLStatements -
var SQLStatements = SQL{
	UpdateEmail:		fmt.Sprintf(updateEmail, constants.Tables.Users),
	UpdatePassword:	fmt.Sprintf(updatePassword, constants.Tables.Users),
	RemoveUser: 		fmt.Sprintf(removeUser, constants.Tables.Users),
	ReviveUser:			fmt.Sprintf(reviveUser, constants.Tables.Users),
}
