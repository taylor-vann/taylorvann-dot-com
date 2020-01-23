// brian taylor vann
// taylorvann dot com

// postgresql statements required of users table

package users

// CreateUsersTableStatement - Create table Statement
const CreateUsersTableStatement = `
CREATE TABLE IF NOT EXISTS users (
  id BIGSERIAL PRIMARY KEY,
  username VARCHAR(64) UNIQUE NOT NULL,
  email VARCHAR(512) UNIQUE NOT NULL,
  is_deleted BOOLEAN NOT NULL DEFAULT FALSE
  created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  updated_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3)
)
`

// InsertUserStatement - Create a user in raw sql
const InsertUserStatement = `
INSERT INTO users (
  username,
  email
)
VALUES (
  $1,
  $2
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