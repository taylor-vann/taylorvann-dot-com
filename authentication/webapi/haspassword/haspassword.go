// brian taylor vann
// taylorvann dot com

// haspassword
// Represents a connection between user and password
// This is a connection, it is also an edge between two vertices

package haspassword

// brian taylor vann
// taylorvann dot com

import (
	"time"
	"webapi/haspassword"
)

// HasPassword - Edge / Connection
type HasPassword struct {
	ID         int64     `json:"id"`
	UserID     int64     `json:"user_id"`
	PasswordID int64     `json:"password_id"`
	CreatedAt  time.Time `json:"created_at"`
}


// CreateParams - arguments for clearer execution
type CreateParams struct {
	UserID     int64
	PasswordID int64
}

// Create - create a password entry in our store
func Create(p *CreateParams) (*PasswordStructure, error) {

}

// Read - update an entry in our store
func Read(p *ReadParams) (*PasswordStructure, error) {
	
}

// Remove - remove an entry from our store
func Remove(p *RemoveParams) (*PasswordStructure, error) {

}