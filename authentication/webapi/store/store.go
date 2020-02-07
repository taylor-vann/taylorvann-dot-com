// Package store - a representation of our database
package store

import "webapi/controllers/users"

// CreateRequiredDatabases - 
func CreateRequiredDatabases() {
	users.CreateTable()
}

// UsersPatchWork - bridge controller to store
type UsersPatchWork struct {}
// Create -
func (*UsersPatchWork) Create(p *users.CreateParams) {
	users.Create(p)
}

// Read -
func (*UsersPatchWork) Read(p *users.ReadParams) {
	users.Read(p)
}

// Update -
func (*UsersPatchWork) Update(p *users.ReadParams, u *users.UpdateParams) {
	users.Update(p, u)
}

// Remove -
func (*UsersPatchWork) Remove(p *users.RemoveParams) {
	users.Remove(p)
}
