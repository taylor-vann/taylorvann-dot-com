// brian taylor vann
// taylorvann dot com
//
// Represents a connection between user and a microservice
//
// all CRUR methods must return entire created or altered entries

// Package roles -  Controller to interact with sql table on device
package roles

import (
	"database/sql"
	"errors"
	"time"

	"webapi/interfaces/storex"
)

// Row -
type Row struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

// CreateParams - arguments for clearer execution
type CreateParams struct {
	UserID int64
	Role   string
}

// ReadParams -
type ReadParams = CreateParams

// RemoveParams -
type RemoveParams = CreateParams

// createUsersRow -
func createRow(rows *sql.Rows) (*Row, error) {
	var rolesRow Row
	if rows.Next() {
		errScan := rows.Scan(
			&rolesRow.ID,
			&rolesRow.UserID,
			&rolesRow.Role,
			&rolesRow.CreatedAt,
		)
		if errScan != nil {
			return nil, errScan
		}
	}

	rows.Close()

	return &rolesRow, nil
}

// CreateTable -
func CreateTable() (*sql.Result, error) {
	result, err := storex.Exec(SQLStatements.CreateTable)
	return &result, err
}

// Create - create a password entry in our store
func Create(p *CreateParams) (*Row, error) {
	if p == nil {
		return nil, errors.New("Nil parameters provided.")
	}

	row, errQueryRow := storex.Query(
		SQLStatements.Create,
		p.UserID,
		p.Role,
	)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return createRow(row)
}

// Read - update an entry in our store
func Read(p *ReadParams) (*Row, error) {
	if p == nil {
		return nil, errors.New("nil parameters provided")
	}

	row, errQueryRow := storex.Query(
		SQLStatements.Read,
		p.UserID,
		p.Role,
	)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return createRow(row)
}

// Remove - remove an entry from our store
func Remove(p *RemoveParams) (*Row, error) {
	if p == nil {
		return nil, errors.New("nil parameters provided")
	}

	row, errQueryRow := storex.Query(
		SQLStatements.Remove,
		p.UserID,
		p.Role,
	)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return createRow(row)
}