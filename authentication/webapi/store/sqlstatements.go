package store

import (
	"fmt"
	"webapi/constants"
)

// DirectStoreSQL -
type DirectStoreSQL struct {
	InsertUserAndPassword string
	RetrieveUserPassword  string
	UpdateUserPassword    string
	RemoveUserAndPassword string
}

const insertUserAndPassword = `
	WITH 
		user_fetch AS (
			SELECT *
			FROM %s
			WHERE
				email = $1
		),
		user_result AS (
			INSERT INTO %s (
				email
			)
			SELECT
				$2
			WHERE
				NOT EXISTS (SELECT * FROM user_fetch)
			RETURNING *
		),
		password_result AS (
			INSERT INTO %s (
				salt,
				hash,
				params
			)
			SELECT
				$3, $4, $5
			WHERE
				NOT EXISTS (SELECT * FROM user_fetch)
			RETURNING
				*
		),
		haspassword_result AS (
			INSERT INTO %s (
				user_id,
				password_id
			)
			SELECT
				(SELECT id FROM user_result),
				(SELECT id FROM password_result)
			WHERE
				NOT EXISTS (SELECT * FROM user_fetch)
			RETURNING
				*
		)
	SELECT * 
	FROM
		user_result
	;
`

const retrieveUserPassword = `
	WITH 
		user_fetch AS (
			SELECT
				*
			FROM 
				%s
			WHERE
				email = $1
		),
		haspassword_result AS (
			SELECT
				*
			FROM
				%s
			WHERE
				EXISTS (SELECT * FROM user_fetch) AND
				user_id = (SELECT id FROM user_fetch)
		),
		password_result AS (
			SELECT
				*
			FROM
				%s
			WHERE
				EXISTS (SELECT * FROM haspassword_result) AND
				id = (SELECT password_id FROM haspassword_result)
		)
	SELECT * 
	FROM
		password_result
	;
`

const updateUserPassword = `
	WITH 
		user_fetch AS (
			SELECT
				*
			FROM 
				%s
			WHERE
				email = $1 AND
				TO_TIMESTAMP($5::DOUBLE PRECISION / 1000) 
					BETWEEN updated_at AND CURRENT_TIMESTAMP(3)
		),
		haspassword_result AS (
			SELECT
				*
			FROM
				%s
			WHERE
				EXISTS (SELECT * FROM user_fetch) AND
				user_id = (SELECT id FROM user_fetch)
		),
		password_update AS (
			UPDATE
				%s
			SET
				salt = $2,
				hash = $3,
				params = $4
			WHERE
				EXISTS (SELECT * FROM haspassword_result) AND
				id = (SELECT password_id FROM haspassword_result)
		)
	SELECT
		* 
	FROM
		user_fetch
	;
`

const removeUserAndPassword = `
	WITH 
		user_removed AS (
			UPDATE
				%s
			SET
				is_deleted = TRUE
			WHERE
				email = $1 AND
				is_deleted = FALSE AND
				TO_TIMESTAMP($2::DOUBLE PRECISION * 0.001) 
					BETWEEN updated_at AND CURRENT_TIMESTAMP(3)
			RETURNING
				*
		),
		haspassword_removed AS (
			DELETE FROM
				%s
			WHERE
				EXISTS (SELECT * FROM user_removed) AND
				user_id = (SELECT id FROM user_removed)
			RETURNING
				*
		),
		password_removed AS (
			DELETE FROM
				%s
			WHERE
				EXISTS (SELECT * FROM haspassword_removed) AND
				id = (SELECT password_id FROM haspassword_removed)
		)
	SELECT
		* 
	FROM
		user_removed
	;
`

func formatInsertUserAndPassword() string {
	return fmt.Sprintf(
		insertUserAndPassword,
		constants.Tables.Users,
		constants.Tables.Users,
		constants.Tables.Passwords,
		constants.Tables.HasPassword,
	)
}

func formatRetrieveUserPassword() string {
	return fmt.Sprintf(
		retrieveUserPassword,
		constants.Tables.Users,
		constants.Tables.HasPassword,
		constants.Tables.Passwords,
	)
}

func formatUpdateUserPassword() string {
	return fmt.Sprintf(
		updateUserPassword,
		constants.Tables.Users,
		constants.Tables.HasPassword,
		constants.Tables.Passwords,
	)
}

func formatRemoveUserAndPassword() string {
	return fmt.Sprintf(
		removeUserAndPassword,
		constants.Tables.Users,
		constants.Tables.HasPassword,
		constants.Tables.Passwords,
	)
}

// SQLStatements -
var SQLStatements = DirectStoreSQL{
	InsertUserAndPassword: formatInsertUserAndPassword(),
	RetrieveUserPassword:  formatRetrieveUserPassword(),
	UpdateUserPassword:    formatUpdateUserPassword(),
	RemoveUserAndPassword: formatRemoveUserAndPassword(),
}
