package responses

import (
	"time"
)

// Sessions (from authn webapi/sessions/hooks/requests)
type Session struct {
	Token string `json:"token"`
}

type ReadEntryParams struct {
	Environment	string `json:"environment"`
	Signature		string `json:"signature"`
}

type SessionErrors struct {
	RequestBody	*string `json:"request_body"`
	Session			*string `json:"session"`
	Default			*string `json:"default"`
}

type SessionBody struct {
	Session				 *Session				 `json:"session"`
	SessionErrors  *SessionErrors  `json:"errors"`
}


// Users (from authn webapi/store/users/hooks/requests)
type SafeRow struct {
	ID        int64       `json:"id"`
	Email     string      `json:"email"`
	IsDeleted bool        `json:"is_deleted"`
	CreatedAt time.Time   `json:"created_at"` // milli seconds
	UpdatedAt time.Time   `json:"updated_at"` // milli seconds
}

type SafeUsers = []SafeRow

type UsersErrors struct {
	RequestBody	*string `json:"request_body"`
	Users 			*string `json:"users"`
	Default 		*string `json:"default"`
}

type UsersBody struct {
	Users  			*SafeUsers 		`json:"users"`
	UsersErrors *UsersErrors	`json:"errors"`
}

type ResponseBodyErrors struct {
	Default 		*string `json:"default"`
	RequestBody	*string `json:"request_body"`
	Users 			*string `json:"users"`
	Session			*string `json:"session"`
}

type ResponseBody struct {
	Errors *ResponseBodyErrors `json:"errors"`
}