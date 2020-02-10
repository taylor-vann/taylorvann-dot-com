// Package passwords
//
// brian taylor vann
// taylorvann dot com
//
// represents a user's password credentials in a store
// it can be considered a vertex
//
// all CRUR methods must return entire created or altered entries

package passwords

import (
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"webapi/interfaces/passwordx"
	"webapi/interfaces/storex"
	"webapi/utils"
)

// HashParams - Expected hash structure
type HashParams = passwordx.HashParams

// PasswordsRow - Expected PostgreSQL structure
type PasswordsRow struct {
	ID        int64       `json:"id"`
	Salt      string      `json:"salt"`
	Hash      string      `json:"hash"`
	Params    *HashParams `json:"params"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

// CreateParams - arguments needed for entry
type CreateParams struct {
	Password string `json:"password"`
}

// ReadParams - arguments needed too remove entry
type ReadParams struct {
	ID int64
}

// UpdateParams - identical arguments needed for password update
type UpdateParams = CreateParams

// RemoveParams - identical arguments needed to remove an entry
type RemoveParams = ReadParams

// CreatePasswordsRow -
func CreatePasswordsRow(rows *sql.Rows) (*PasswordsRow, error) {
	var password PasswordsRow
	if rows.Next() {
		var jsonParamsAsStr string
		errScan := rows.Scan(
			&password.ID,
			&password.Salt,
			&password.Hash,
			&jsonParamsAsStr,
			&password.UpdatedAt,
			&password.CreatedAt,
		)
		if errScan != nil {
			return nil, errScan
		}

		errMarshal := json.Unmarshal([]byte(jsonParamsAsStr), &password.Params)
		if errMarshal != nil {
			return nil, errMarshal
		}

		return &password, nil
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
func Create(p *CreateParams) (*PasswordsRow, error) {
	if p == nil {
		return nil, errors.New("passwords:Create - nil parameters")
	}
	hashedPassword, errHashPassword := passwordx.HashPassword(
		p.Password,
		&passwordx.DefaultHashParams,
	)
	if errHashPassword != nil {
		return nil, errHashPassword
	}

	marshaledParams, errMarshal := json.Marshal(passwordx.DefaultHashParams)
	if errMarshal != nil {
		return nil, errMarshal
	}

	row, errQueryRow := storex.Query(
		SQLStatements.Create,
		hashedPassword.Salt,
		hashedPassword.Hash,
		marshaledParams,
	)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return CreatePasswordsRow(row)
}

// Read - update an entry in our store
func Read(p *ReadParams) (*PasswordsRow, error) {
	if p == nil {
		return nil, errors.New("passwords:Read - nil parameters")
	}

	row, errQueryRow := storex.Query(SQLStatements.Read, p.ID)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return CreatePasswordsRow(row)
}

// Update - update an entry in our store
func Update(readParams *ReadParams, updateParams *UpdateParams) (*PasswordsRow, error) {
	if readParams == nil {
		return nil, errors.New("passwords:Update - nil readParams")
	}
	if updateParams == nil {
		return nil, errors.New("passwords:Update - nil updateParams")
	}

	hashedPassword, errHashPassword := passwordx.HashPassword(
		updateParams.Password,
		&passwordx.DefaultHashParams,
	)
	if errHashPassword != nil {
		return nil, errHashPassword
	}

	marshaledParams, errMarshal := json.Marshal(passwordx.DefaultHashParams)
	if errMarshal != nil {
		return nil, errMarshal
	}

	row, errQueryRow := storex.Query(
		SQLStatements.Update,
		readParams.ID,
		hashedPassword.Salt,
		hashedPassword.Hash,
		marshaledParams,
		utils.GetNowAsMS(),
	)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return CreatePasswordsRow(row)
}

// Remove - remove an entry from our store
func Remove(p *RemoveParams) (*PasswordsRow, error) {
	row, errQueryRow := storex.Query(
		SQLStatements.Remove,
		p.ID,
		utils.GetNowAsMS(),
	)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return CreatePasswordsRow(row)
}
