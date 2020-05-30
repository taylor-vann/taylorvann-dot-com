package responses

import "time"

// responses
type Session struct {
	Token string
}

type Errors struct {
	Headers			*string `json:"headers"`
	RequestBody	*string `json:"request_body"`
	Session			*string `json:"session"`
	Default			*string `json:"default"`
}

type Body struct {
	Session *Session	`json:"session"`
	Errors  *Errors		`json:"errors"`
}

type User struct {
	ID        int64       `json:"id"`
	Email     string      `json:"email"`
	IsDeleted bool        `json:"is_deleted"`
	CreatedAt time.Time   `json:"created_at"` // milli seconds
	UpdatedAt time.Time   `json:"updated_at"` // milli seconds
}

type UserErrors struct {
	Headers			*string `json:"headers"`
	RequestBody	*string `json:"request_body"`
	Session			*string `json:"session"`
	Default			*string `json:"default"`
}

type Users = []User

type UsersBody struct {
	Users 	*Users				`json:"users"`
	Errors  *UserErrors		`json:"errors"`
}