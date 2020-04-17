// brian taylor vann
// taylorvann dot com

package roles

import (
	"database/sql"
	"errors"
	"time"

	"webapi/interfaces/storex"
	"webapi/controllers/roles/constants"
	"webapi/controllers/roles/statements"
	"webapi/controllers/roles/utils"
)

type Row struct {
	ID					 int64     `json:"id"`
	UserID    	 int64     `json:"user_id"`
	Organization string    `json:"organization"`
	ReadAccess	 bool			 `json:"read_access"`
	WriteAccess	 bool			 `json:"write_access"`
	IsDeleted		 bool			 `json:"is_deleted"`
	CreatedAt		 time.Time `json:"created_at"`
	UpdatedAt		 time.Time `json:"updated_at"`
}

type Roles = []Row

type CreateTableParams struct {
	Environment string `json:"environment"`
}

type CreateParams struct {
	Environment  string
	UserID			 int64
	Organization string
	ReadAccess	 bool
	WriteAccess	 bool
}

type ReadParams struct {
	Environment  string
	UserID			 int64
	Organization string
}

type IndexParams struct {
	Environment  string
	StartIndex	 int64
	Length  		 int64
}

type SearchParams struct {
	Environment string
	UserID 			int64
}

type UpdateParams struct {
	Environment  string
	UserID			 int64
	Organization string
	ReadAccess	 bool
	WriteAccess	 bool
	IsDeleted		 bool
}

type UpdateAccessParams = CreateParams
type DeleteParams = ReadParams
type UndeleteParams = ReadParams

func getDefaultEnvironment(environment string) string {
	if environment != "" {
		return environment
	}
	return constants.RolesTest
}

func CreateTable(p *CreateTableParams) (*sql.Result, error) {
	environment := getDefaultEnvironment(p.Environment)
	statement := statements.SqlMap[environment].CreateTable

	result, err := storex.Exec(statement)
	return &result, err
}

func CreateRows(rows *sql.Rows) (Roles, error) {
	if rows == nil {
		return nil, errors.New("roles.CreateRows() - nil params provided")
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
		return nil, errors.New("Nil parameters provided.")
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
		return nil, errors.New("nil parameters provided")
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

func Search(p *SearchParams) (Roles, error) {
	if p == nil {
		return nil, errors.New("nil parameters provided")
	}

	environment := getDefaultEnvironment(p.Environment)
	statement := statements.SqlMap[environment].Search

	rows, errQueryRows := storex.Query(
		statement,
		p.UserID,
	)
	if errQueryRows != nil {
		return nil, errQueryRows
	}

	return CreateRows(rows)
}

func Update(p *UpdateParams) (Roles, error) {
	if p == nil {
		return nil, errors.New("nil parameters provided")
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
		return nil, errors.New("nil parameters provided")
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
		return nil, errors.New("nil parameters provided")
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
		return nil, errors.New("nil parameters provided")
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
