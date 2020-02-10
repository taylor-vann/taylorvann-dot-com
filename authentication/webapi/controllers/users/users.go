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
	"errors"
	"time"

	"webapi/interfaces/storex"
	"webapi/utils"
)

// UsersRow - Expected PostgreSQL structure
type UsersRow struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	IsDeleted bool      `json:"is_deleted"`
	CreatedAt time.Time `json:"created_at"` // milli seconds
	UpdatedAt time.Time `json:"updated_at"` // milli seconds
}

// Users -
type Users = []UsersRow

// CreateParams - arguments needed for entry
type CreateParams struct {
	Email string `json:"email"`
}

// ReadParams - arguments needed too remove entry
type ReadParams = CreateParams

// UpdateParams - identical arguments needed for password update
type UpdateParams struct {
	CurrentEmail       string
	UpdatedEmail       string
	RequestedTimestamp int64
}

// RemoveParams - identical arguments needed to remove an entry
type RemoveParams = CreateParams

// CreateUsersRow -
func CreateUsersRow(rows *sql.Rows) (*UsersRow, error) {
	var user UsersRow
	if rows.Next() {
		errScan := rows.Scan(
			&user.ID,
			&user.Email,
			&user.IsDeleted,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if errScan != nil {
			return nil, errScan
		}

		return &user, nil
	}

	rows.Close()

	return nil, nil
}

// CreateTable -
func CreateTable() (*sql.Result, error) {
	result, err := storex.Exec(SQLStatements.CreateTable)
	return &result, err
}

// Create - create a password entry in our store
func Create(p *CreateParams) (*UsersRow, error) {
	if p == nil {
		return nil, errors.New("nil parameters provided")
	}

	row, errQueryRow := storex.Query(
		SQLStatements.Create,
		p.Email,
	)

	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return CreateUsersRow(row)
}

// Read - update an entry in our store
func Read(p *ReadParams) (*UsersRow, error) {
	if p == nil {
		return nil, errors.New("nil parameters provided")
	}

	row, errQueryRow := storex.Query(
		SQLStatements.Read,
		p.Email,
	)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return CreateUsersRow(row)
}

// Update - update an entry in our store
func Update(p *UpdateParams) (*UsersRow, error) {
	if p == nil {
		return nil, errors.New("nil parameters provided")
	}

	row, errQueryRow := storex.Query(
		SQLStatements.Update,
		p.CurrentEmail,
		p.UpdatedEmail,
		utils.GetNowAsMS(),
	)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return CreateUsersRow(row)
}

// Remove - remove an entry from our store
func Remove(p *RemoveParams) (*UsersRow, error) {
	if p == nil {
		return nil, errors.New("nil parameters provided")
	}

	row, errQueryRow := storex.Query(
		SQLStatements.Remove,
		p.Email,
		utils.GetNowAsMS(),
	)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return CreateUsersRow(row)
}
