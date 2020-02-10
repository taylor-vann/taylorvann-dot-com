// brian taylor vann
// taylorvann dot com

// haspassword
// Represents a connection between user and password
// This is a connection, it is also an edge between two vertices

package haspassword

// brian taylor vann
// taylorvann dot com

import (
	"database/sql"
	"errors"
	"time"

	"webapi/interfaces/storex"
	"webapi/utils"
)

// HasPasswordRow -
type HasPasswordRow struct {
	ID         int64     `json:"id"`
	UserID     int64     `json:"user_id"`
	PasswordID int64     `json:"password_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// CreateParams - arguments for clearer execution
type CreateParams struct {
	UserID     int64
	PasswordID int64
}

type ReadParams struct {
	UserID int64
}

type UpdateParams = CreateParams

type RemoveParams = ReadParams

// createUsersRow -
func createHasPasswordRow(rows *sql.Rows) (*HasPasswordRow, error) {
	var haspassword HasPasswordRow
	if rows.Next() {
		errScan := rows.Scan(
			&haspassword.ID,
			&haspassword.UserID,
			&haspassword.PasswordID,
			&haspassword.CreatedAt,
			&haspassword.UpdatedAt,
		)
		if errScan != nil {
			return nil, errScan
		}
	}

	rows.Close()

	return &haspassword, nil
}

// CreateTable -
func CreateTable() (*sql.Result, error) {
	result, err := storex.Exec(SQLStatements.CreateTable)
	return &result, err
}

// Create - create a password entry in our store
func Create(p *CreateParams) (*HasPasswordRow, error) {
	if p == nil {
		return nil, errors.New("Nil parameters provided.")
	}

	row, errQueryRow := storex.Query(
		SQLStatements.Create,
		p.UserID,
		p.PasswordID,
	)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return createHasPasswordRow(row)
}

// Read - update an entry in our store
func Read(p *ReadParams) (*HasPasswordRow, error) {
	if p == nil {
		return nil, errors.New("nil parameters provided")
	}

	row, errQueryRow := storex.Query(SQLStatements.Read, p.UserID)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return createHasPasswordRow(row)
}

// Update - update an entry in our store
func Update(p *UpdateParams) (*HasPasswordRow, error) {
	if p == nil {
		return nil, errors.New("nil parameters provided")
	}

	row, errQueryRow := storex.Query(
		SQLStatements.Update,
		p.UserID,
		p.PasswordID,
		utils.GetNowAsMS(),
	)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return createHasPasswordRow(row)
}

// Remove - remove an entry from our store
func Remove(p *RemoveParams) (*HasPasswordRow, error) {
	if p == nil {
		return nil, errors.New("nil parameters provided")
	}

	row, errQueryRow := storex.Query(
		SQLStatements.Remove,
		p.UserID,
		utils.GetNowAsMS(),
	)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return createHasPasswordRow(row)
}
