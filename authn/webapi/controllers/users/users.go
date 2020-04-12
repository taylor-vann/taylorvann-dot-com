// brian taylor vann
// taylorvann dot com
//

package users

import (
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"webapi/controllers/utils"
	"webapi/interfaces/passwordx"
	"webapi/interfaces/storex"
)

// HashParams -
type HashParams = passwordx.HashParams

// Row - Expected PostgreSQL structure
type Row struct {
	ID        int64       `json:"id"`
	Email     string      `json:"email"`
	IsDeleted bool        `json:"is_deleted"`
	Salt      string      `json:"salt"`
	Hash      string      `json:"hash"`
	Params    *HashParams `json:"params"`
	CreatedAt time.Time   `json:"created_at"` // milli seconds
	UpdatedAt time.Time   `json:"updated_at"` // milli seconds
}

// Users -
type Users = []Row

// CreateParams - arguments needed for entry
type CreateParams struct {
	Email    string	`json:"email"`	
	Password string	`json:"password"`
}

// ReadParams - arguments needed too remove entry
type ReadParams struct {
	Email string	`json:"email"`
}

// SearchParams - arguments needed too remove entry
type SearchParams struct {
	EmailSubstring string	`json:"email_substring"`
}

// UpdateParams - identical arguments needed for password update
type UpdateParams struct {
	CurrentEmail  string	`json:"current_email"`
	UpdatedEmail  string	`json:"updated_email"`
	Password      string	`json:"password"`
	IsDeleted			bool		`json:"is_deleted"`
}

// UpdateEmailParams - identical arguments needed for password update
type UpdateEmailParams struct {
	CurrentEmail  string	`json:"current_email"`
	UpdatedEmail  string	`json:"updated_email"`
}

// UpdatePasswordParams - identical arguments needed for password update
type UpdatePasswordParams = CreateParams

// RemoveParams - identical arguments needed to remove an entry
type RemoveParams = ReadParams

// ReviveParams - identical arguments needed to remove an entry
type ReviveParams = ReadParams


// CreateRows -
func CreateRows(rows *sql.Rows) (Users, error) {
	if rows == nil {
		return nil, errors.New("users.CreateRows() - nil params provided")
	}

	var users Users

	defer rows.Close()
	for rows.Next() {
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

		users = append(users, userRow)
	}

	return users, nil
}

// CreateTable -
func CreateTable() (*sql.Result, error) {
	result, err := storex.Exec(SQLStatements.CreateTable)
	return &result, err
}

// Create - create a password entry in our store
func Create(p *CreateParams) (Users, error) {
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

	rows, errQueryRow := storex.Query(
		SQLStatements.Create,
		p.Email,
		hashedPassword.Salt,
		hashedPassword.Hash,
		marshaledParams,
	)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return CreateRows(rows)
}

// Read - update an entry in our store
func Read(p *ReadParams) (Users, error) {
	if p == nil {
		return nil, errors.New("users.Read() - nil parameters provided")
	}

	rows, errQueryRow := storex.Query(
		SQLStatements.Read,
		p.Email,
	)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return CreateRows(rows)
}

// Search - find entries similar to a substring
func Search(p *SearchParams) (Users, error) {
	if p == nil {
		return nil, errors.New("users.Search() - nil parameters provided")
	}

	rows, errQueryRow := storex.Query(
		SQLStatements.Search,
		p.EmailSubstring,
	)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return CreateRows(rows)
}

// Update - update an entry in our store
func Update(p *UpdateParams) (Users, error) {
	if p == nil {
		return nil, errors.New("users.Updated() - nil parameters provided")
	}

	if p.Password == "" {
		return nil, errors.New("users.Updated() - password cannot be empty string")
	}

	if p.CurrentEmail == "" {
		return nil, errors.New("users.Updated() - current email cannot be empty string")
	}

	if p.UpdatedEmail == "" {
		return nil, errors.New("users.Updated() - updated email cannot be empty string")
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

	rows, errQueryRow := storex.Query(
		SQLStatements.Update,
		p.CurrentEmail,
		p.UpdatedEmail,
		p.IsDeleted,
		hashedPassword.Salt,
		hashedPassword.Hash,
		marshaledParams,
		utils.GetNowAsMS(),
	)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return CreateRows(rows)
}

// UpdateEmail - update an entry email in our store
func UpdateEmail(p *UpdateEmailParams) (Users, error) {
	if p == nil {
		return nil, errors.New("users.Updated() - nil parameters provided")
	}

	if p.CurrentEmail == "" {
		return nil, errors.New("users.Updated() - current email cannot be empty string")
	}

	if p.UpdatedEmail == "" {
		return nil, errors.New("users.Updated() - updated email cannot be empty string")
	}

	rows, errQueryRow := storex.Query(
		SQLStatements.UpdateEmail,
		p.CurrentEmail,
		p.UpdatedEmail,
		utils.GetNowAsMS(),
	)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return CreateRows(rows)
}

// Update - update an entry password in our store
func UpdatePassword(p *UpdatePasswordParams) (Users, error) {
	if p == nil {
		return nil, errors.New("users.Updated() - nil parameters provided")
	}

	if p.Password == "" {
		return nil, errors.New("users.Updated() - password cannot be empty string")
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

	rows, errQueryRow := storex.Query(
		SQLStatements.UpdatePassword,
		p.Email,
		hashedPassword.Salt,
		hashedPassword.Hash,
		marshaledParams,
		utils.GetNowAsMS(),
	)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return CreateRows(rows)
}

// Remove - remove an entry from our store
func Remove(p *RemoveParams) (Users, error) {
	if p == nil {
		return nil, errors.New("users.Remove() - nil parameters provided")
	}

	rows, errQueryRow := storex.Query(
		SQLStatements.Remove,
		p.Email,
		utils.GetNowAsMS(),
	)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return CreateRows(rows)
}

// Revive - revive an entry from our store
func Revive(p *RemoveParams) (Users, error) {
	if p == nil {
		return nil, errors.New("users.Revive() - nil parameters provided")
	}

	rows, errQueryRow := storex.Query(
		SQLStatements.Revive,
		p.Email,
		utils.GetNowAsMS(),
	)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return CreateRows(rows)
}
