// brian taylor vann
// taylorvann dot com

// passwords
// represents a user's password credentials
// it can be considered a vertex

package passwords

import (
	"time"
	"webapi/interfaces/passwordx"
)

// PasswordHashParams - Expected hash structure
type PasswordHashParams = passwordx.HashParams

// PasswordStructure - Expected PostgreSQL structure
type PasswordStructure struct {
	ID        int64              `json:"id"`
	UserID    int64              `json:"user_id"`
	Salt      string             `json:"salt"`
	Hash      string             `json:"hash"`
	Params    PasswordHashParams `json:"params"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}
