package requests

type UserCredentials struct {
	UserID int64 `json:"user_id"`
}

type AccountCredentials struct {
	Email		 string	`json:"email"`
	Password string	`json:"password"`
}

type Params struct {
	Environment				 string 						 `json:"environment"`
	SessionToken			 string			 		 		 `json:"session_token"`
	AccountCredentials *AccountCredentials `json:"account_credentials"`
	UserCredentials		 *UserCredentials		 `json:"user_credentials"`
}

type Body struct {
	Action string  `json:"action"`
	Params *Params `json:"params"`
}
