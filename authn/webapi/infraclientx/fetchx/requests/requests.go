package requests

type GuestSession struct {
	Environment string  `json:"environment"`
}

type ValidateSession struct {
	Environment string  `json:"environment"`
	Token				string	`json:"token"`
}

type ValidateGuestUser struct {
	Environment string `json:"environment"`
	Email				string `json:"email"`
	Password		string `json:"password"`
}

type ValidateRoleFromSession struct {
	Environment  string `json:"environment`
	Token				 string	`json:"token"`
	Organization string `json:"organization"`
}

type ValidateRole struct {
	Environment  string `json:"environment`
	UserID			 int64	`json:"user_id"`
	Organization string `json:"organization"`
}

type ValidateInfraRole = ValidateGuestUser
type InfraSession = ValidateGuestUser
type ValidateUser = ValidateGuestUser


type Body struct {
	Action string 		 `json:"action"`
	Params interface{} `json:"params"`
}