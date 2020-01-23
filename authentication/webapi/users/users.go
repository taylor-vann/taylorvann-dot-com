package users

// UserStructure - Expected structure from PostgreSQL
type UserStructure struct {
	id int `json:"id"`
	username  string `json:"username"`
	createdAt int `json:"created_at"`
	updatedAt int `json:"updated_at"`
	isDeleted string `json:"is_deleted"`
}
