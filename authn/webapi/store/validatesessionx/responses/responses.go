package responses

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