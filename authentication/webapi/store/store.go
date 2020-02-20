// brian taylor vann
// taylorvann dot com

// All methods (except ValidateUser) return a user row.
// We cache queries. We cache mutations. They all return user rows.

// Package store - a representation of our database
package store

import (
	"encoding/json"
	"errors"
	"time"

	"webapi/controllers/users"
	"webapi/interfaces/passwordx"
	"webapi/interfaces/storex"
)

// MilliSeconds -
type MilliSeconds = int64

// CreateUserParams -
type CreateUserParams = users.CreateParams

// ReadUserParams -
type ReadUserParams = users.ReadParams

// UpdateEmailParams -
type UpdateEmailParams struct {
	CurrentEmail string
	UpdatedEmail string
}

// ValidateUserParams -
type ValidateUserParams struct {
	Email    string
	Password string
}

// RemoveUserParams -
type RemoveUserParams = ReadUserParams

// UpdatePasswordParams -
type UpdatePasswordParams struct {
	Email           string
	UpdatedPassword string
}

// UserRow -
type UserRow = users.Row

// getNowAsMS -
func getNowAsMS() MilliSeconds {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// CreateRequiredTables -
func CreateRequiredTables() (bool, error) {
	_, errUsers := users.CreateTable()
	if errUsers != nil {
		return false, errors.New("Error creating users table")
	}

	return true, nil
}

// CreateUser -
func CreateUser(p *CreateUserParams) (*UserRow, error) {
	return users.Create(p)
}

// ReadUser -
func ReadUser(p *ReadUserParams) (*UserRow, error) {
	return users.Read(p)
}

// ValidateUser - creates a new person in our store, returns UserRow
func ValidateUser(p *ValidateUserParams) (*UserRow, error) {
	if p == nil {
		return nil, errors.New("store.ValidateUser() - nil parameters given")
	}
	userRow, errUserRow := users.Read(&ReadUserParams{
		Email: p.Email,
	})

	if errUserRow != nil {
		return nil, errUserRow
	}

	hashResults := passwordx.HashResults{
		Salt:   userRow.Salt,
		Hash:   userRow.Hash,
		Params: userRow.Params,
	}

	result, errPasswordVaildation := passwordx.PasswordIsValid(
		p.Password,
		&hashResults,
	)

	if errPasswordVaildation != nil {
		return nil, errPasswordVaildation
	}

	if result == true {
		return userRow, nil
	}

	return nil, nil
}

// UpdateEmail - creates a new person in our store, returns UserRow
func UpdateEmail(p *UpdateEmailParams) (*UserRow, error) {
	if p == nil {
		return nil, errors.New("store.UpdateEmail() - nil parameters given")
	}
	userRow, errUserRow := storex.Query(
		SQLStatements.UpdateEmail,
		p.CurrentEmail,
		p.UpdatedEmail,
		getNowAsMS(),
	)
	if errUserRow != nil {
		return nil, errUserRow
	}

	return users.CreateRow(userRow)
}

// UpdatePassword - creates a new person in our store, returns UserRow
func UpdatePassword(p *UpdatePasswordParams) (*UserRow, error) {
	if p == nil {
		return nil, errors.New("store.UpdatePassword() - nil parameters given")
	}
	hashResults, errHashResults := passwordx.HashPassword(
		p.UpdatedPassword,
		&passwordx.DefaultHashParams,
	)
	if errHashResults != nil {
		return nil, errHashResults
	}

	marshaledParams, errMarshal := json.Marshal(hashResults.Params)
	if errMarshal != nil {
		return nil, errMarshal
	}

	userRow, errUserRow := storex.Query(
		SQLStatements.UpdatePassword,
		p.Email,
		hashResults.Salt,
		hashResults.Hash,
		marshaledParams,
		getNowAsMS(),
	)

	if errUserRow != nil {
		return nil, errUserRow
	}

	return users.CreateRow(userRow)
}

// RemoveUser - creates a new person in our store, returns UserRow
func RemoveUser(p *RemoveUserParams) (*UserRow, error) {
	return users.Remove(p)
}
