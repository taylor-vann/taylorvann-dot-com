package users

import (
	"webapi/users"
)

// CreateUsersTableStatement - Create table Statement
const CreateUsersTableStatement = `
CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  username varchar(64) UNIQUE NOT NULL,
  email varchar(512) UNIQUE NOT NULL,
  created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  updated_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  is_deleted BOOLEAN NOT NULL DEFAULT FALSE
)
`

// InsertUserStatement - Create a user in raw sql
const InsertUserStatement = `
INSERT INTO users (
  username,
  email,
)
VALUES (
  $1,
  $2,
RETURNING
	*;
`

// UpdateUserStatement - Update a users email or password
const UpdateUserStatement = `
UPDATE
	users
SET
  email = $2,
  username = $3,
  updated_at = CURRENT_TIMESTAMP(3)
WHERE
	id = $1
RETURNING 
	*;
`

// UpdateUserAsDeletedStatement - Soft delete our user
const UpdateUserAsDeletedStatement = `
UPDATE
  users
SET
  is_deleted = TRUE,
  updated_at = CURRENT_TIMESTAMP(3)
WHERE
	id = $1
RETURNING
	*;
`