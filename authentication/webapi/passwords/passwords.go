// Package passwords
//
// brian taylor vann
// taylorvann dot com
//
// represents a user's password credentials in a store
// it can be considered a vertex
//
// all CRUR methods must return entire created or altered entries

package passwords

import (
	"errors"
	"time"

	"webapi/interfaces/passwordx"
)

// PasswordHashParams - Expected hash structure
type PasswordHashParams = passwordx.HashParams

// PasswordStructure - Expected PostgreSQL structure
type PasswordStructure struct {
	ID        int64              `json:"id"`
	Salt      string             `json:"salt"`
	Hash      string             `json:"hash"`
	Params    PasswordHashParams `json:"params"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

// CreateParams - arguments needed for entry 
type CreateParams struct {
	Salt      string             `json:"salt"`
	Hash      string             `json:"hash"`
	Params    PasswordHashParams `json:"params"`
}

// ReadParams - arguments needed too remove entry
type ReadParams struct {
	id int
}

// UpdateParams - identical arguments needed for password update
type UpdateParams = CreateParams

// RemoveParams - identical arguments needed to remove an entry
type RemoveParams = ReadParams

// Create - create a password entry in our store
func Create(p *CreateParams) (*PasswordStructure, error) {

}

// Read - update an entry in our store
func Read(p *ReadParams) (*PasswordStructure, error) {
	
}

// Update - update an entry in our store
func Update(p *UpdateParams) (*PasswordStructure, error) {
	
}

// Remove - remove an entry from our store
func Remove(p *RemoveParams) (*PasswordStructure, error) {

}