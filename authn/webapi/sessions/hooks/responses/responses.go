package responses

type SessionPayload struct {
	SessionToken	string	`json:"session_token"`
	CsrfToken			string	`json:"csrf_token"`
}

type ErrorsPayload struct {
	Headers	*string `json:"headers"`
	Body		*string `json:"body"`
	Session *string `json:"session"`
	Default *string `json:"default"`
}

type Body struct {
	Session *SessionPayload	`json:"session"`
	Errors  *ErrorsPayload	`json:"errors"`
}
