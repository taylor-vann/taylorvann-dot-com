// brian taylor vann
// taylorvann dot com

package controller

import (
	"testing"
)

var createTable = CreateTableParams{
	Environment: "LOCAL",
}

var testRolesCreate = CreateParams{
	Environment: "LOCAL",
	UserID: 1,
	Organization: "internal_admin",
	ReadAccess: true,
	WriteAccess: true,
}

var testRolesCreate2 = CreateParams{
	Environment: "LOCAL",
	UserID: 1,
	Organization: "internal_rebel",
	ReadAccess: true,
	WriteAccess: true,
}

var testRolesRead = ReadParams{
	Environment: "LOCAL",
	UserID: 1,
	Organization:   "internal_admin",
}

var testRolesSearch = SearchParams{
	Environment: "LOCAL",
	UserID: 1,
	StartIndex: 0,
	Length: 10,
}

var testRolesIndex = IndexParams{
	Environment: "LOCAL",
	StartIndex: 0,
	Length: 1024,
}

var testRolesUpdate = UpdateParams{
	Environment: "LOCAL",
	UserID: 1,
	Organization: "internal_admin",
	ReadAccess: false,
	WriteAccess: false,
	IsDeleted: false,
}

var testRolesUpdateAccess = UpdateAccessParams{
	Environment: "LOCAL",
	UserID: 1,
	Organization: "internal_rebel",
	ReadAccess: false,
	WriteAccess: false,
}

func TestCreateTable(t *testing.T) {
	results, err := CreateTable(&createTable)
	if err != nil {
		t.Error(err.Error())
	}
	if results == nil {
		t.Error("no results were returned from CreateTable.")
	}
}

func TestCreate(t *testing.T) {
	rows, err := Create(&testRolesCreate)
	if err != nil {
		t.Error(err.Error())
	}
	if len(rows) == 0 {
		t.Error("No results were returned from Create.")
		return
	}
	if len(rows) != 1 {
		t.Error("Incorrect amount of results were returned from Create.")
		return
	}

	result := rows[0]

	if result.UserID != testRolesCreate.UserID {
		t.Error("failed to create Roles row.")
	}
	if result.Organization != testRolesCreate.Organization {
		t.Error("failed to create Roles row.")
	}
}

func TestRead(t *testing.T) {
	rows, err := Read(&testRolesRead)
	if err != nil {
		t.Error(err.Error())
	}
	if len(rows) == 0 {
		t.Error("No results were returned from Create.")
		return
	}
	if len(rows) != 1 {
		t.Error("Incorrect amount of results were returned from Create.")
		return
	}

	result := rows[0]

	if result.UserID != testRolesCreate.UserID {
		t.Error("mutated UserID.")
	}
	if result.Organization != testRolesCreate.Organization {
		t.Error("failed to read roles.")
	}
}

func TestSearch(t *testing.T) {
	rowsCreate, errCreate := Create(&testRolesCreate2)
	if errCreate != nil {
		t.Error(errCreate.Error())
	}
	if len(rowsCreate) == 0 {
		t.Error("No results were returned from Create.")
		return
	}
	if len(rowsCreate) != 1 {
		t.Error("Incorrect amount of results were returned from Create.")
		return
	}

	rows, err := Search(&testRolesSearch)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if len(rows) == 0 {
		t.Error("No results were returned from Search.")
		return
	}

	if len(rows) < 2 {
		t.Error("More than two results were supposed to be returned")
		return
	}
}

func TestIndex(t *testing.T) {
	rows, err := Index(&testRolesIndex)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if len(rows) == 0 {
		t.Error("No results were returned from Index.")
		return
	}

	if len(rows) < 2 {
		t.Error("More than two results were supposed to be returned")
		return
	}
}

func TestUpdate(t *testing.T) {
	rows, err := Update(&testRolesUpdate)
	if err != nil {
		t.Error(err.Error())
	}
	if len(rows) == 0 {
		t.Error("No results were returned from Update.")
		return
	}
	if len(rows) != 1 {
		t.Error("Incorrect amount of results were returned from Update.")
		return
	}

	result := rows[0]

	if result.UserID != testRolesCreate.UserID {
		t.Error("Mutated UserID.")
	}
	if result.ReadAccess != false {
		t.Error("Role read access is incorrect.")
	}
	if result.WriteAccess != false {
		t.Error("Role write access is incorrect.")
	}
}

func TestUpdateAccess(t *testing.T) {
	rows, err := UpdateAccess(&testRolesUpdateAccess)
	if err != nil {
		t.Error(err.Error())
	}
	if len(rows) == 0 {
		t.Error("No results were returned from Update.")
		return
	}
	if len(rows) != 1 {
		t.Error("Incorrect amount of results were returned from Update.")
		return
	}

	result := rows[0]

	if result.UserID != testRolesCreate2.UserID {
		t.Error("Mutated UserID.")
	}
	if result.ReadAccess != false {
		t.Error("Role read access is incorrect.")
	}
	if result.WriteAccess != false {
		t.Error("Role write access is incorrect.")
	}
}

func TestDelete(t *testing.T) {
	rows, err := Delete(&testRolesRead)
	if err != nil {
		t.Error(err.Error())
	}
	if len(rows) == 0 {
		t.Error("No results were returned from Delete.")
		return
	}
	if len(rows) != 1 {
		t.Error("Incorrect amount of results were returned from Delete.")
		return
	}

	result := rows[0]

	if result.UserID != testRolesCreate.UserID {
		t.Error("Mutated UserID.")
	}
	if result.IsDeleted != true {
		t.Error("Role has not been deleted.")
	}
}

func TestUndelete(t *testing.T) {
	rows, err := Undelete(&testRolesRead)
	if err != nil {
		t.Error(err.Error())
	}
	if len(rows) == 0 {
		t.Error("No results were returned from Delete.")
		return
	}
	if len(rows) != 1 {
		t.Error("Incorrect amount of results were returned from Delete.")
		return
	}

	result := rows[0]

	if result.UserID != testRolesCreate.UserID {
		t.Error("Mutated UserID.")
	}
	if result.IsDeleted != false {
		t.Error("Role has not been undeleted.")
	}
}

func TestDangerouslyDropUnitTestsTable(t *testing.T) {
	result, err := DangerouslyDropUnitTestsTable()
	if result == nil {
		t.Error("Failed to drop table")
	}
	if err != nil {
		t.Error(err.Error())
	}
}