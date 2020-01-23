// brian taylor vann
// taylorvann dot com

// haspassword
// Represents a connection between user and password
// This is a connection, it is also an edge between two vertices

package haspassword

// brian taylor vann
// taylorvann dot com

import "time"

// HasPassword - Edge / Connection
type HasPassword struct {
	ID         int64     `json:"id"`
	UserID     int64     `json:"user_id"`
	PasswordID int64     `json:"password_id"`
	CreatedAt  time.Time `json:"created_at"`
}
