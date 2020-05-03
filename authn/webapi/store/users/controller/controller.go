// brian taylor vann
// taylorvann dot com

package controller

import (
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"webapi/interfaces/passwordx"
	"webapi/interfaces/storex"
	"webapi/store/users/controller/constants"
	"webapi/store/users/controller/statements"
	"webapi/store/users/controller/utils"
)

type HashParams = passwordx.HashParams

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

type Users = []Row

type CreateTableParams struct {
	Environment string `json:"environment"`
}

type CreateParams struct {
	Environment string `json:"environment"`
	Email    		string `json:"email"`	
	Password 		string `json:"password"`
}

type ReadParams struct {
	Environment string `json:"environment"`
	Email 			string `json:"email"`
}

type IndexParams struct {
	Environment  string `json:"environment"`
	StartIndex	 int64  `json:"start_index"`
	Length  		 int64	`json:"length"`
}

type SearchParams struct {
	Environment 	 string `json:"environment"`
	EmailSubstring string	`json:"email_substring"`
	StartIndex	 int64  `json:"start_index"`
	Length  		 int64	`json:"length"`
}

type UpdateParams struct {
	Environment 	string `json:"environment"`
	CurrentEmail  string `json:"current_email"`
	UpdatedEmail  string `json:"updated_email"`
	Password      string `json:"password"`
	IsDeleted			bool	 `json:"is_deleted"`
}

type UpdateEmailParams struct {
	Environment 	string `json:"environment"`
	CurrentEmail  string `json:"current_email"`
	UpdatedEmail  string `json:"updated_email"`
}

type UpdatePasswordParams = CreateParams
type DeleteParams = ReadParams
type UndeleteParams = ReadParams


func getDefaultEnvironment(environment string) string {
	if environment != "" {
		return environment
	}
	return constants.UsersTest
}

func CreateTable(p *CreateTableParams) (*sql.Result, error) {
	environment := getDefaultEnvironment(p.Environment)
	statement := statements.SqlMap[environment].CreateTable

	result, err := storex.Exec(statement)
	return &result, err
}

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

func Create(p *CreateParams) (Users, error) {
	if p == nil {
		return nil, errors.New("nil parameters provided")
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

	environment := getDefaultEnvironment(p.Environment)
	statement := statements.SqlMap[environment].Create
	rows, errQueryRows := storex.Query(
		statement,
		p.Email,
		hashedPassword.Salt,
		hashedPassword.Hash,
		marshaledParams,
	)

	if errQueryRows != nil {
		return nil, errQueryRows
	}

	return CreateRows(rows)
}

func Read(p *ReadParams) (Users, error) {
	if p == nil {
		return nil, errors.New("nil parameters provided")
	}

	environment := getDefaultEnvironment(p.Environment)
	statement := statements.SqlMap[environment].Read
	rows, errQueryRow := storex.Query(
		statement,
		p.Email,
	)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return CreateRows(rows)
}

func Index(p *IndexParams) (Users, error) {
	if p == nil {
		return nil, errors.New("nil parameters provided")
	}

	environment := getDefaultEnvironment(p.Environment)
	statement := statements.SqlMap[environment].Index

	rows, errQueryRows := storex.Query(
		statement,
		p.StartIndex,
		p.Length,
	)
	if errQueryRows != nil {
		return nil, errQueryRows
	}

	return CreateRows(rows)
}

func Search(p *SearchParams) (Users, error) {
	if p == nil {
		return nil, errors.New("nil parameters provided")
	}

	environment := getDefaultEnvironment(p.Environment)
	statement := statements.SqlMap[environment].Search
	rows, errQueryRow := storex.Query(
		statement,
		p.EmailSubstring,
	)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return CreateRows(rows)
}

func Update(p *UpdateParams) (Users, error) {
	if p == nil {
		return nil, errors.New("nil parameters provided")
	}

	if p.Password == "" {
		return nil, errors.New("password cannot be empty string")
	}

	if p.CurrentEmail == "" {
		return nil, errors.New("current email cannot be empty string")
	}

	if p.UpdatedEmail == "" {
		return nil, errors.New("updated email cannot be empty string")
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

	environment := getDefaultEnvironment(p.Environment)
	statement := statements.SqlMap[environment].Update
	rows, errQueryRow := storex.Query(
		statement,
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

func UpdateEmail(p *UpdateEmailParams) (Users, error) {
	if p == nil {
		return nil, errors.New("nil parameters provided")
	}

	if p.CurrentEmail == "" {
		return nil, errors.New("current email cannot be empty string")
	}

	if p.UpdatedEmail == "" {
		return nil, errors.New("updated email cannot be empty string")
	}

	environment := getDefaultEnvironment(p.Environment)
	statement := statements.SqlMap[environment].UpdateEmail
	rows, errQueryRow := storex.Query(
		statement,
		p.CurrentEmail,
		p.UpdatedEmail,
		utils.GetNowAsMS(),
	)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return CreateRows(rows)
}

func UpdatePassword(p *UpdatePasswordParams) (Users, error) {
	if p == nil {
		return nil, errors.New("nil parameters provided")
	}

	if p.Password == "" {
		return nil, errors.New("password cannot be empty string")
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

	environment := getDefaultEnvironment(p.Environment)
	statement := statements.SqlMap[environment].UpdatePassword
	rows, errQueryRow := storex.Query(
		statement,
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

func Delete(p *DeleteParams) (Users, error) {
	if p == nil {
		return nil, errors.New("nil parameters provided")
	}

	environment := getDefaultEnvironment(p.Environment)
	statement := statements.SqlMap[environment].Delete
	rows, errQueryRow := storex.Query(
		statement,
		p.Email,
		utils.GetNowAsMS(),
	)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return CreateRows(rows)
}

func Undelete(p *UndeleteParams) (Users, error) {
	if p == nil {
		return nil, errors.New("nil parameters provided")
	}

	environment := getDefaultEnvironment(p.Environment)
	statement := statements.SqlMap[environment].Undelete
	rows, errQueryRow := storex.Query(
		statement,
		p.Email,
		utils.GetNowAsMS(),
	)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return CreateRows(rows)
}

func DangerouslyDropUnitTestsTable() (sql.Result, error) {
	return storex.Exec(statements.DangerouslyDropUnitTestsTable)
}
