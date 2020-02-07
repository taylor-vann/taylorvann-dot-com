// Package Users
//
// brian taylor vann
// taylorvann dot com
//
// represents a user's password credentials in a store
// it can be considered a vertex
//
// all CRUR methods must return entire created or altered entries

package users

import (
	"database/sql"
	"time"

	"webapi/interfaces/storex"
)

// UsersRow - Expected PostgreSQL structure
type UsersRow struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	IsDeleted bool      `json:"is_deleted"`
	CreatedAt time.Time `json:"created_at"` // milli seconds
	UpdatedAt time.Time `json:"updated_at"` // milli seconds
}

// CreateParams - arguments needed for entry
type CreateParams struct {
	Email string `json:"email"`
}

// ReadParams - arguments needed too remove entry
type ReadParams = CreateParams

// UpdateParams - identical arguments needed for password update
type UpdateParams = CreateParams

// RemoveParams - identical arguments needed to remove an entry
type RemoveParams = CreateParams

// createUsersRow -
func createUsersRow(row *sql.Row) (*UsersRow, error) {
	var user UsersRow
	errScan := row.Scan(
		&user.ID,
		&user.Email,
		&user.IsDeleted,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	return &user, errScan
}

// CreateTable -
func CreateTable() (sql.Result, error) {
	return storex.Exec(SQLStatements.CreateTable)
}

// Create - create a password entry in our store
func Create(p *CreateParams) (*UsersRow, error) {
	row, errQueryRow := storex.QueryRow(SQLStatements.Create, p.Email)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return createUsersRow(row)
}

// Read - update an entry in our store
func Read(p *ReadParams) (*UsersRow, error) {
	row, errQueryRow := storex.QueryRow(SQLStatements.Read, p.Email)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return createUsersRow(row)
}

// Update - update an entry in our store
func Update(readParams *ReadParams, updateParams *UpdateParams) (*UsersRow, error) {
	row, errQueryRow := storex.QueryRow(
		SQLStatements.Update,
		readParams.Email,
		updateParams.Email,
	)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return createUsersRow(row)
}

// Remove - remove an entry from our store
func Remove(p *RemoveParams) (*UsersRow, error) {
	row, errQueryRow := storex.QueryRow(SQLStatements.Remove, p.Email)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return createUsersRow(row)
}
