// Package store - a representation of our database
package store

import (
	"encoding/json"
	"fmt"

	"webapi/controllers/passwords"
	"webapi/controllers/users"
	"webapi/interfaces/passwordx"
	"webapi/interfaces/storex"
	"webapi/utils"
)

// CreateUserParams -
type CreateUserParams struct {
	Email    string
	Username string
	Password string
}

// ReadUserParams -
type ReadUserParams = users.ReadParams

// UpdateUserParams -
type UpdateUserParams = users.UpdateParams

// ValidateUserParams -
type ValidateUserParams struct {
	Email    string
	Password string
}

// RemoveUserParams -
type RemoveUserParams = ReadUserParams

// UpdateUserPasswordParams -
type UpdateUserPasswordParams struct {
	Email           string
	UpdatedPassword string
}

// UserRow -
type UserRow = users.UsersRow

// InsertUserAndPasswordParams -
type InsertUserAndPasswordParams struct {
	email            string
	salt             string
	hash             string
	params           string
	requestTimestamp int64 // milli seconds
}

// CreateUser -
func CreateUser(p *CreateUserParams) (*UserRow, error) {
	hashResults, errHashPassword := passwordx.HashPassword(
		p.Password,
		&passwordx.DefaultHashParams,
	)
	if errHashPassword != nil {
		fmt.Println(errHashPassword)
		return nil, errHashPassword
	}

	marshaledParams, errMarshal := json.Marshal(hashResults.Params)
	if errMarshal != nil {
		return nil, errMarshal
	}

	userRow, errUserRow := storex.Query(
		SQLStatements.InsertUserAndPassword,
		p.Email,
		p.Email,
		hashResults.Salt,
		hashResults.Hash,
		marshaledParams,
	)
	if errUserRow != nil {
		fmt.Println(errUserRow)
		return nil, errUserRow
	}

	user, errUser := users.CreateUsersRow(userRow)
	if errUser != nil {
		return nil, errUser
	}

	return user, nil
}

// ReadUser -
func ReadUser(p *ReadUserParams) (*UserRow, error) {
	return users.Read(p)
}

// UpdateUser - creates a new person in our store, returns UserRow
func UpdateUser(p *UpdateUserParams) (*UserRow, error) {
	return users.Update(p)
}

// ValidateUser - creates a new person in our store, returns UserRow
func ValidateUser(p *ValidateUserParams) (bool, error) {
	passwordRow, errPasswordRow := storex.Query(
		SQLStatements.RetrieveUserPassword,
		p.Email,
	)

	if errPasswordRow != nil {
		return false, errPasswordRow
	}

	password, errPassword := passwords.CreatePasswordsRow(passwordRow)
	if errPassword != nil {
		return false, errPassword
	}

	if password == nil {
		return false, nil
	}

	hashResults := passwordx.HashResults{
		Salt:   password.Salt,
		Hash:   password.Hash,
		Params: password.Params,
	}

	result, errPasswordVaildation := passwordx.PasswordIsValid(
		p.Password,
		&hashResults,
	)

	if errPasswordVaildation != nil {
		return false, errPasswordVaildation
	}

	return result, nil
}

// UpdateUserPassword - creates a new person in our store, returns UserRow
func UpdateUserPassword(p *UpdateUserPasswordParams) (*UserRow, error) {
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
		SQLStatements.UpdateUserPassword,
		p.Email,
		hashResults.Salt,
		hashResults.Hash,
		marshaledParams,
		utils.GetNowAsMS(),
	)

	if errUserRow != nil {
		return nil, errUserRow
	}

	user, errUser := users.CreateUsersRow(userRow)
	if errUser != nil {
		return nil, errUser
	}

	return user, nil
}

// RemoveUser - creates a new person in our store, returns UserRow
func RemoveUser(p *RemoveUserParams) (*UserRow, error) {
	userRow, errUserRow := storex.Query(
		SQLStatements.RemoveUserAndPassword,
		p.Email,
		utils.GetNowAsMS(),
	)

	if errUserRow != nil {
		return nil, errUserRow
	}

	user, errUser := users.CreateUsersRow(userRow)
	if errUser != nil {
		return nil, errUser
	}

	return user, nil
}
