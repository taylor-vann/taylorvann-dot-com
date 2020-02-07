package store

import "webapi/controllers/users"

type UsersRow = users.UsersRow

// CreateParams -
type UsersCreateParams = users.CreateParams

// ReadParams -
type UsersReadParams = users.ReadParams

// UpdateParams -
type UsersUpdateParams = users.UpdateParams

// RemoveParams -
type UsersRemoveParams = users.RemoveParams

// UsersStore - A cheap way to interface with our users controller
type UsersStore struct{}

// Every store requires:
// - Create
// - Read
// - Update
// - Remove

// Create - 
func (*UsersStore) Create(p *UsersCreateParams) (*UsersRow, error) {
	return users.Create(p)
}

// Read -
func (*UsersStore) Read(p *UsersReadParams) (*UsersRow, error) {
	return users.Read(p)
}

// Update -
func (*UsersStore) Update(p *UsersReadParams, u *UsersUpdateParams) (*UsersRow, error) {
	return users.Update(p, u)
}

// Remove -
func (*UsersStore) Remove(p *UsersRemoveParams) (*UsersRow, error) {
	return users.Remove(p)
}

// Users -
var Users = UsersPatchWork{}
