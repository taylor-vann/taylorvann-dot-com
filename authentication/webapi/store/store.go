// Package store - a representation of our database
package store

import (
	"fmt"
	"webapi/interfaces/storex"
	"webapi/statements"
)


// CreateRequiredDatabases - generate tables in our store
func CreateRequiredDatabases() {
	fmt.Println("CreateRequiredDatabases")
	fmt.Println(statements.Users.CreateTable)

	var result, errUserTable = storex.Exec(statements.Users.CreateTable)

	fmt.Println(result)
	fmt.Println(errUserTable)

	storex.Exec(statements.UsersTest.CreateTable)

	storex.Exec(statements.Passwords.CreateTable)
	storex.Exec(statements.PasswordsTest.CreateTable)

	storex.Exec(statements.HasPasswords.CreateTable)
	storex.Exec(statements.HasPasswordsTest.CreateTable)
}
