package responses

type UserCredentials struct {
	UserID int64 `json:"user_id"`
}

type AccountCredentials struct {
	Email		 string	`json:"email"`
	Password string	`json:"password"`
}

type RequestParams struct {
	Environment				 string 						 `json:"environment"`
	SessionToken			 string			 		 		 `json:"session_token"`
	AccountCredentials *AccountCredentials `json:"account_credentials"`
	UserCredentials		 *UserCredentials		 `json:"user_credentials"`
}

type RequestBody struct {
	Action string         `json:"action"`
	Params *RequestParams `json:"params"`
}

type SessionResponsePayload struct {
	SessionToken	string	`json:"session_token"`
	CsrfToken			string	`json:"csrf_token"`
}

type ErrorsResponsePayload struct {
	Headers	*string `json:"headers"`
	Body		*string `json:"body"`
	Session *string `json:"session"`
	Default *string `json:"default"`
}

type ResponseBody struct {
	Session *SessionResponsePayload	`json:"session"`
	Errors  *ErrorsResponsePayload	`json:"errors"`
}
