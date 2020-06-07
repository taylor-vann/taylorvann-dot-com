package responses

import "time"

// responses
type Session struct {
	Token string
}

type SessionErrors struct {
	Headers			*string `json:"headers"`
	RequestBody	*string `json:"request_body"`
	Session			*string `json:"session"`
	Default			*string `json:"default"`
}

type SessionBody struct {
	Session *Session				`json:"session"`
	Errors  *SessionErrors	`json:"errors"`
}

type User struct {
	ID        int64       `json:"id"`
	Email     string      `json:"email"`
	IsDeleted bool        `json:"is_deleted"`
	CreatedAt time.Time   `json:"created_at"` // milli seconds
	UpdatedAt time.Time   `json:"updated_at"` // milli seconds
}

type UserErrors struct {
	RequestBody	*string `json:"request_body"`
	Users				*string `json:"users"`
	Default			*string `json:"default"`
}

type Users = []User

type UsersBody struct {
	Users 	*Users				`json:"users"`
	Errors  *UserErrors		`json:"errors"`
}

type Role struct {
	ID					 int64     `json:"id"`
	UserID    	 int64     `json:"user_id"`
	Organization string    `json:"organization"`
	ReadAccess	 bool			 `json:"read_access"`
	WriteAccess	 bool			 `json:"write_access"`
	IsDeleted		 bool			 `json:"is_deleted"`
	CreatedAt		 time.Time `json:"created_at"`
	UpdatedAt		 time.Time `json:"updated_at"`
}

type Roles = []Role

type RolesErrors struct {
	RequestBody	*string `json:"request_body"`
	Roles				*string `json:"roles"`
	Default			*string `json:"default"`
}

type RolesBody struct {
	Roles 	*Roles				`json:"roles"`
	Errors  *RolesErrors	`json:"errors"`
}