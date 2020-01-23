// brian taylor vann
// taylorvann dot com

package users

import (
	"time"
)

// UserStructure - Expected structure from PostgreSQL
type UserStructure struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	IsDeleted bool      `json:"is_deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
