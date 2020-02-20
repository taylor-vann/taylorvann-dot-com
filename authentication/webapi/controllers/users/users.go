// brian taylor vann
// taylorvann dot com
//
// represents a user's password credentials in a store
// it can be considered a vertex
//
// all CRUR methods must return entire created or altered entries

// Package users - Controller to interact with sql table on device
package users

import (
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"webapi/interfaces/passwordx"
	"webapi/interfaces/storex"
	"webapi/utils"
)

// HashParams -
type HashParams = passwordx.HashParams

// Row - Expected PostgreSQL structure
type Row struct {
	ID        int64       `json:"id"`
	Email     string      `json:"email"`
	Salt      string      `json:"salt"`
	Hash      string      `json:"hash"`
	Params    *HashParams `json:"params"`
	IsDeleted bool        `json:"is_deleted"`
	CreatedAt time.Time   `json:"created_at"` // milli seconds
	UpdatedAt time.Time   `json:"updated_at"` // milli seconds
}

// Users -
type Users = []Row

// CreateParams - arguments needed for entry
type CreateParams struct {
	Email  		string 
	Password 	string
}

// ReadParams - arguments needed too remove entry
type ReadParams struct {
	Email string
}

// UpdateParams - identical arguments needed for password update
type UpdateParams struct {
	CurrentEmail				string
	UpdatedEmail				string
	Password						string
	RequestedTimestamp	int64
}

// RemoveParams - identical arguments needed to remove an entry
type RemoveParams = ReadParams

// CreateRow -
func CreateRow(rows *sql.Rows) (*Row, error) {
	if rows == nil {
		return nil, errors.New("users.CreateRow() - nil params provided")
	}
	rows.Next()
	var userRow Row
	var jsonParamsAsStr string
	errScan := rows.Scan(
		&userRow.ID,
		&userRow.Email,
		&userRow.Salt,
		&userRow.Hash,
		&jsonParamsAsStr,
		&userRow.IsDeleted,
		&userRow.CreatedAt,
		&userRow.UpdatedAt,
	)
	rows.Close()
	if errScan != nil {
		return nil, errScan
	}

	errMarshal := json.Unmarshal(
		[]byte(jsonParamsAsStr),
		&userRow.Params,
	)
	if errMarshal != nil {
		return nil, errMarshal
	}

	return &userRow, nil
}

// CreateTable -
func CreateTable() (*sql.Result, error) {
	result, err := storex.Exec(SQLStatements.CreateTable)
	return &result, err
}

// Create - create a password entry in our store
func Create(p *CreateParams) (*Row, error) {
	if p == nil {
		return nil, errors.New("users.Create() - nil parameters provided")
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
		p.Email,
		hashedPassword.Salt,
		hashedPassword.Hash,
		marshaledParams,
	)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return CreateRow(row)
}

// Read - update an entry in our store
func Read(p *ReadParams) (*Row, error) {
	if p == nil {
		return nil, errors.New("users.Read() - nil parameters provided")
	}

	row, errQueryRow := storex.Query(
		SQLStatements.Read,
		p.Email,
	)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return CreateRow(row)
}

// Update - update an entry in our store
func Update(p *UpdateParams) (*Row, error) {
	if p == nil {
		return nil, errors.New("users.Updated() - nil parameters provided")
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
		SQLStatements.Update,
		p.CurrentEmail,
		p.UpdatedEmail,
		hashedPassword.Salt,
		hashedPassword.Hash,
		marshaledParams,
		utils.GetNowAsMS(),
	)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return CreateRow(row)
}

// Remove - remove an entry from our store
func Remove(p *RemoveParams) (*Row, error) {
	if p == nil {
		return nil, errors.New("users.Remove() - nil parameters provided")
	}

	row, errQueryRow := storex.Query(
		SQLStatements.Remove,
		p.Email,
		utils.GetNowAsMS(),
	)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return CreateRow(row)
}
