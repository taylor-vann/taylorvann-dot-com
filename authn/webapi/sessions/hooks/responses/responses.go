package responses

type Session struct {
	SessionToken	string	`json:"session_token"`
	CsrfToken			string	`json:"csrf_token"`
}

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
