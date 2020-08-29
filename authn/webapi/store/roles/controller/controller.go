// brian taylor vann
// briantaylorvann dot com

package controller

import (
	"database/sql"
	"errors"
	"time"

	"webapi/store/roles/controller/constants"
	"webapi/store/roles/controller/statements"
	"webapi/store/roles/controller/utils"

	"github.com/taylor-vann/weblog/toolbox/golang/storex"
)

type Row struct {
	ID           int64     `json:"id"`
	UserID       int64     `json:"user_id"`
	Organization string    `json:"organization"`
	ReadAccess   bool      `json:"read_access"`
	WriteAccess  bool      `json:"write_access"`
	IsDeleted    bool      `json:"is_deleted"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Roles = []Row

type CreateTableParams struct {
	Environment string `json:"environment"`
}

type CreateParams struct {
	Environment  string `json:"environment`
	UserID       int64  `json:"user_id"`
	Organization string `json:"organization"`
	ReadAccess   bool   `json:"read_access"`
	WriteAccess  bool   `json:"write_access"`
}

type ReadParams struct {
	Environment  string `json:"environment`
	UserID       int64  `json:"user_id"`
	Organization string `json:"organization"`
}

type IndexParams struct {
	Environment string `json:"environment`
	StartIndex  int64  `json:"start_index"`
	Length      int64  `json:"length"`
}

type SearchParams struct {
	Environment string `json:"environment`
	UserID      int64  `json:"user_id"`
	StartIndex  int64  `json::"start_index"`
	Length      int64  `json:"length"`
}

type UpdateParams struct {
	Environment  string `json:"environment`
	UserID       int64  `json:"user_id"`
	Organization string `json:"organization"`
	ReadAccess   bool   `json:"read_access"`
	WriteAccess  bool   `json:"write_access"`
	IsDeleted    bool   `json:"is_deleted"`
}

type UpdateAccessParams = CreateParams
type DeleteParams = ReadParams
type UndeleteParams = ReadParams

var (
	errNilParams = errors.New("nil params provided")
)

func getDefaultEnvironment(environment string) string {
	if environment != "" {
		return environment
	}

	if constants.Environment == constants.Development {
		return constants.Development
	}

	return constants.Local
}

func CreateTable(p *CreateTableParams) (*sql.Result, error) {
	environment := getDefaultEnvironment(p.Environment)
	statement := statements.SqlMap[environment].CreateTable

	result, err := storex.Exec(statement)
	return &result, err
}

func CreateRows(rows *sql.Rows) (Roles, error) {
	if rows == nil {
		return nil, errNilParams
	}

	var roles Roles

	defer rows.Close()
	for rows.Next() {
		var rolesRow Row

		errScan := rows.Scan(
			&rolesRow.ID,
			&rolesRow.UserID,
			&rolesRow.Organization,
			&rolesRow.ReadAccess,
			&rolesRow.WriteAccess,
			&rolesRow.IsDeleted,
			&rolesRow.CreatedAt,
			&rolesRow.UpdatedAt,
		)
		if errScan != nil {
			return nil, errScan
		}

		roles = append(roles, rolesRow)
	}

	return roles, nil
}

func Create(p *CreateParams) (Roles, error) {
	if p == nil {
		return nil, errNilParams
	}

	environment := getDefaultEnvironment(p.Environment)
	statement := statements.SqlMap[environment].Create

	rows, errQueryRows := storex.Query(
		statement,
		p.UserID,
		p.Organization,
		p.ReadAccess,
		p.WriteAccess,
	)

	if errQueryRows != nil {
		return nil, errQueryRows
	}

	return CreateRows(rows)
}

func Read(p *ReadParams) (Roles, error) {
	if p == nil {
		return nil, errNilParams
	}

	environment := getDefaultEnvironment(p.Environment)
	statement := statements.SqlMap[environment].Read

	rows, errQueryRows := storex.Query(
		statement,
		p.UserID,
		p.Organization,
	)
	if errQueryRows != nil {
		return nil, errQueryRows
	}

	return CreateRows(rows)
}

func Index(p *IndexParams) (Roles, error) {
	if p == nil {
		return nil, errNilParams
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

func Search(p *SearchParams) (Roles, error) {
	if p == nil {
		return nil, errNilParams
	}

	environment := getDefaultEnvironment(p.Environment)
	statement := statements.SqlMap[environment].Search

	rows, errQueryRows := storex.Query(
		statement,
		p.UserID,
		p.StartIndex,
		p.Length,
	)
	if errQueryRows != nil {
		return nil, errQueryRows
	}

	return CreateRows(rows)
}

func Update(p *UpdateParams) (Roles, error) {
	if p == nil {
		return nil, errNilParams
	}

	environment := getDefaultEnvironment(p.Environment)
	statement := statements.SqlMap[environment].Update

	rows, errQueryRows := storex.Query(
		statement,
		p.UserID,
		p.Organization,
		p.ReadAccess,
		p.WriteAccess,
		p.IsDeleted,
		utils.GetNowAsMS(),
	)
	if errQueryRows != nil {
		return nil, errQueryRows
	}

	return CreateRows(rows)
}

func UpdateAccess(p *UpdateAccessParams) (Roles, error) {
	if p == nil {
		return nil, errNilParams
	}

	environment := getDefaultEnvironment(p.Environment)
	statement := statements.SqlMap[environment].UpdateAccess

	rows, errQueryRows := storex.Query(
		statement,
		p.UserID,
		p.Organization,
		p.ReadAccess,
		p.WriteAccess,
		utils.GetNowAsMS(),
	)
	if errQueryRows != nil {
		return nil, errQueryRows
	}

	return CreateRows(rows)
}

func Delete(p *DeleteParams) (Roles, error) {
	if p == nil {
		return nil, errNilParams
	}

	environment := getDefaultEnvironment(p.Environment)
	statement := statements.SqlMap[environment].Delete

	row, errQueryRow := storex.Query(
		statement,
		p.UserID,
		p.Organization,
		utils.GetNowAsMS(),
	)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return CreateRows(row)
}

func Undelete(p *UndeleteParams) (Roles, error) {
	if p == nil {
		return nil, errNilParams
	}

	environment := getDefaultEnvironment(p.Environment)
	statement := statements.SqlMap[environment].Undelete

	row, errQueryRow := storex.Query(
		statement,
		p.UserID,
		p.Organization,
		utils.GetNowAsMS(),
	)
	if errQueryRow != nil {
		return nil, errQueryRow
	}

	return CreateRows(row)
}

func DangerouslyDropUnitTestsTable() (sql.Result, error) {
	return storex.Exec(statements.DangerouslyDropUnitTestsTable)
}
