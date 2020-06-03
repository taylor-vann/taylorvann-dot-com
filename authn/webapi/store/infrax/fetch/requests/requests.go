package requests

type GuestSessionParams struct {
	Environment string  `json:"environment"`
}

type ValidateGuestSessionParams = GuestSessionParams

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