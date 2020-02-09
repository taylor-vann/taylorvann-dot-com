// Package store - a representation of our database
package store

import (
	"encoding/json"
	"fmt"
	"time"

	"webapi/controllers/haspassword"
	"webapi/controllers/passwords"
	"webapi/controllers/users"
	"webapi/interfaces/passwordx"
	"webapi/interfaces/storex"
)

// CreateUserParams -
type CreateUserParams struct {
	Email    string
	Username string
	Password string
}

// UpdateUserParams -
type UpdateUserParams struct {
	Email        string
	Password     string
	UpdatedEmail string
}

// ValidateUserParams -
type ValidateUserParams struct {
	Email    string
	Password string
}

// RemoveUserParams -
type RemoveUserParams = CreateUserParams

// UpdatePasswordParams -
type UpdatePasswordParams struct {
	Email           string
	Password        string
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

// MilliSecond -
type MilliSecond = int64

// User exists
// Username exists

// Write Operations
// CreateUser - creates a new person in our store, returns UserRow

func getNowAsMS() MilliSecond {
	return time.Now().UnixNano() / int64(time.Millisecond)
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

	// nowAsMS := getNowAsMS()

	marshaledParams, errMarshal := json.Marshal(hashResults.Params)
	if errMarshal != nil {
		return nil, errMarshal
	}

	userRow, errUserRow := storex.QueryRow(
		SQLStatements.InsertUserAndPassword,
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

// // UpdateUser - creates a new person in our store, returns UserRow
// func UpdateUser(p *UpdateUserParams) (*UserRow, error) {
// 	// update user

// 	// return user
// }

// // UpdateUserPassword - creates a new person in our store, returns UserRow
// func UpdateUserPassword(p *UpdatePasswordParams) (*UserRow, error) {
// 	// get user
// 	// get haspassword
// 	// get password

// 	// return user
// }

// ValidateUser - creates a new person in our store, returns UserRow
func ValidateUser(p *ValidateUserParams) (bool, error) {
	// get user
	userParams := users.ReadParams{
		Email: p.Email,
	}
	userRow, errUserRow := users.Read(&userParams)
	if errUserRow != nil {
		return false, errUserRow
	}

	// get haspassword
	haspasswordParams := haspassword.ReadParams{
		UserID: userRow.ID,
	}
	haspasswordRow, errHasPasswordRow := haspassword.Read(
		&haspasswordParams,
	)
	if errHasPasswordRow != nil {
		return false, errHasPasswordRow
	}

	// get password
	passwordParams := passwords.ReadParams{
		ID: haspasswordRow.PasswordID,
	}
	passwordRow, errPasswordRow := passwords.Read(&passwordParams)
	if errPasswordRow != nil {
		return false, errPasswordRow
	}

	// check result
	hashResults := passwordx.HashResults{
		Salt:   passwordRow.Salt,
		Hash:   passwordRow.Hash,
		Params: passwordRow.Params,
	}
	result, errPasswordVaildation := passwordx.PasswordIsValid(
		p.Password,
		&hashResults,
	)

	fmt.Println(hashResults)
	fmt.Println(p.Password)
	if errPasswordVaildation != nil {
		return false, errPasswordVaildation
	}

	return result, nil
}

// RemoveUser - creates a new person in our store, returns UserRow
func RemoveUser(p *RemoveUserParams) (*UserRow, error) {
	// Remove user
	userParams := users.RemoveParams{
		Email: p.Email,
	}
	userRow, errUserRow := users.Remove(&userParams)
	if errUserRow != nil {
		return nil, errUserRow
	}

	// Remove haspassword as link
	haspasswordParams := haspassword.RemoveParams{
		UserID: userRow.ID,
	}
	haspasswordRow, errHasPasswordRow := haspassword.Remove(
		&haspasswordParams,
	)
	if errHasPasswordRow != nil {
		return nil, errHasPasswordRow
	}

	// Remove password
	passwordParams := passwords.RemoveParams{
		ID: haspasswordRow.PasswordID,
	}
	_, errPasswordRow := passwords.Remove(&passwordParams)
	if errPasswordRow != nil {
		return nil, errPasswordRow
	}

	// return user
	return userRow, nil
}
