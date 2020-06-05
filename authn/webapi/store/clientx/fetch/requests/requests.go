package requests

type GuestSession struct {
	Environment string  `json:"environment"`
}

type ValidateGuestSession = GuestSession
type ValidateSession struct {
	Environment string  `json:"environment"`
	Token				string	`json:"token"`
}
type ValidateGuestUser struct {
	Environment string `json:"environment"`
	Email				string `json:"email"`
	Password		string `json:"password"`
}

type ValidateInfraRole = ValidateGuestUser
type InfraSession = ValidateGuestUser

type Body struct {
	Action string 		 `json:"action"`
	Params interface{} `json:"params"`
}