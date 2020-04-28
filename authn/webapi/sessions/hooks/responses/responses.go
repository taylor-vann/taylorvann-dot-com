package responses

import (
	"webapi/sessions/sessionsx"
)

type Session = sessionsx.Session

type Errors struct {
	Headers	*string `json:"headers"`
	Body		*string `json:"body"`
	Session *string `json:"session"`
	Default *string `json:"default"`
}

type Body struct {
	Session *Session	`json:"session"`
	Errors  *Errors		`json:"errors"`
}
